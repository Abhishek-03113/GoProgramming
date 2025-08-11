package main

import (
	"fmt"
	"time"
)

type FileType int

const (
	Game FileType = iota
	Zip
	Movie
	Document
)

func Download(fileType FileType, downloadChannel chan int) {

	start := time.Now()
	for progress := int(fileType) + 1; progress <= 100; progress *= 2 {
		downloadChannel <- progress
		if int(fileType) == 0 {
			time.Sleep(time.Duration(progress) * time.Millisecond)
		}
		time.Sleep(time.Second)

	}

	duration := time.Since(start)
	fmt.Printf("Downloaded file of type %+v, in %v \n", fileType, duration)
	close(downloadChannel)
}

func main() {

	gameDownloadChan := make(chan int)
	zipDownloadChan := make(chan int)
	movieDownlodChan := make(chan int)
	documentDownloadChan := make(chan int)

	go Download(Game, gameDownloadChan)
	go Download(Zip, zipDownloadChan)
	go Download(Movie, movieDownlodChan)
	go Download(Document, documentDownloadChan)

	for {
		select {
		case movieDownloadProgress, ok := <-movieDownlodChan:

			if !ok {
				fmt.Println("Movie Download Complete 👍")
				movieDownlodChan = nil
			} else {
				fmt.Printf("Movie Download Progress := %d \n", movieDownloadProgress)

			}
		case zipDownloadProgress, ok := <-zipDownloadChan:

			if !ok {
				fmt.Println("Zip Download Complete 👍")
				zipDownloadChan = nil
			} else {
				fmt.Printf("zip Download Progress := %d \n", zipDownloadProgress)
			}
		case gameDownloadProgress, ok := <-gameDownloadChan:

			if !ok {
				fmt.Println("Game Download Complete 👍")
				gameDownloadChan = nil
			} else {
				fmt.Printf("game Download Progress := %d \n", gameDownloadProgress)
			}

		case documentDownloadProgress, ok := <-documentDownloadChan:

			if !ok {
				fmt.Println("Document Download Complete 👍 ")
				documentDownloadChan = nil
			} else {
				fmt.Printf("documet Download Progress := %d \n", documentDownloadProgress)
			}

		case <-time.After(100 * time.Second):
			fmt.Println("Download Speed Very slow exitting ")
			return

		}

		if gameDownloadChan == nil && movieDownlodChan == nil && zipDownloadChan == nil && documentDownloadChan == nil {
			fmt.Println("All files Downloaded Successfully, !! Exiting the loop")
			break
		}
	}

}
