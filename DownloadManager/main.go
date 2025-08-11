package main

import (
	"fmt"
	"sync"
	"time"
)

type FileType int

const (
	Game FileType = iota
	Zip
	Movie
	Document
)

func Download(fileType FileType, downloadChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(downloadChannel)
	for progress := int(fileType) + 1; progress <= 100; progress *= 2 {
		downloadChannel <- progress
		if int(fileType) == 0 {
			time.Sleep(time.Duration(progress) * time.Millisecond)
		}
		time.Sleep(time.Second)
	}
}

func PrintProgress(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for progress := range ch {
		fmt.Printf("%s progress: %d%%\n", name, progress)
	}
	eta := time.Since(start)
	fmt.Printf("%s complete 👍 in %v\n", name, eta)
}

func main() {
	var wg sync.WaitGroup
	var printWg sync.WaitGroup

	gameDownloadChan := make(chan int, 1)
	zipDownloadChan := make(chan int, 1)
	movieDownloadChan := make(chan int, 1)
	documentDownloadChan := make(chan int, 1)

	wg.Add(4)
	go Download(Game, gameDownloadChan, &wg)
	go Download(Zip, zipDownloadChan, &wg)
	go Download(Movie, movieDownloadChan, &wg)
	go Download(Document, documentDownloadChan, &wg)

	printWg.Add(4)
	go PrintProgress("Game", gameDownloadChan, &printWg)
	go PrintProgress("Zip", zipDownloadChan, &printWg)
	go PrintProgress("Movie", movieDownloadChan, &printWg)
	go PrintProgress("Document", documentDownloadChan, &printWg)

	wg.Wait()
	printWg.Wait()

	fmt.Println("All files downloaded successfully!")
}
