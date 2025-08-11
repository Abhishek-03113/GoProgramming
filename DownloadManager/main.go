//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//type FileType int
//
//const (
//	Game FileType = iota
//	Zip
//	Movie
//	Document
//)
//
//func Download(fileType FileType, downloadChannel chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	defer close(downloadChannel)
//	start := time.Now()
//	for progress := int(fileType) + 1; progress <= 100; progress *= 2 {
//		downloadChannel <- progress
//		if int(fileType) == 0 {
//			time.Sleep(time.Duration(progress) * time.Millisecond)
//		}
//		time.Sleep(time.Second)
//
//	}
//
//	duration := time.Since(start)
//	fmt.Printf("Downloaded file of type %+v, in %v \n", fileType, duration)
//}

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

func Download(fileType FileType, downloadChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(downloadChannel)

	start := time.Now()
	for progress := int(fileType) + 1; progress <= 100; progress *= 2 {
		downloadChannel <- progress
		if int(fileType) == 0 {
			time.Sleep(time.Duration(progress) * time.Millisecond)
		}
		time.Sleep(time.Second)
	}
	duration := time.Since(start)
	fmt.Printf("Downloaded file of type %+v in %v\n", fileType, duration)
}

func PrintProgress(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for progress := range ch {
		fmt.Printf("%s progress: %d%%\n", name, progress)
	}
	fmt.Printf("%s complete 👍\n", name)
}

func main() {
	var wg sync.WaitGroup
	var printWg sync.WaitGroup

	gameDownloadChan := make(chan int, 1)
	zipDownloadChan := make(chan int, 1)
	movieDownloadChan := make(chan int, 1)
	documentDownloadChan := make(chan int, 1)

	// Start download goroutines
	wg.Add(4)
	go Download(Game, gameDownloadChan, &wg)
	go Download(Zip, zipDownloadChan, &wg)
	go Download(Movie, movieDownloadChan, &wg)
	go Download(Document, documentDownloadChan, &wg)

	// Start progress printing goroutines
	printWg.Add(4)
	go PrintProgress("Game", gameDownloadChan, &printWg)
	go PrintProgress("Zip", zipDownloadChan, &printWg)
	go PrintProgress("Movie", movieDownloadChan, &printWg)
	go PrintProgress("Document", documentDownloadChan, &printWg)

	// Wait for downloads to finish
	wg.Wait()
	// Wait for printing to finish
	printWg.Wait()

	fmt.Println("All files downloaded successfully!")
}

//
//func main() {
//	var wg sync.WaitGroup
//
//	gameDownloadChan := make(chan int, 1)
//	zipDownloadChan := make(chan int, 1)
//	movieDownlodChan := make(chan int, 1)
//	documentDownloadChan := make(chan int, 1)
//
//	wg.Add(1)
//	go Download(Game, gameDownloadChan, &wg)
//	wg.Add(1)
//	go Download(Zip, zipDownloadChan, &wg)
//	wg.Add(1)
//	go Download(Movie, movieDownlodChan, &wg)
//	wg.Add(1)
//	go Download(Document, documentDownloadChan, &wg)
//
//	for {
//		select {
//		case movieDownloadProgress, ok := <-movieDownlodChan:
//
//			if !ok {
//				fmt.Println("Movie Download Complete 👍")
//				movieDownlodChan = nil
//			} else {
//				fmt.Printf("Movie Download Progress := %d \n", movieDownloadProgress)
//
//			}
//		case zipDownloadProgress, ok := <-zipDownloadChan:
//
//			if !ok {
//				fmt.Println("Zip Download Complete 👍")
//				zipDownloadChan = nil
//			} else {
//				fmt.Printf("zip Download Progress := %d \n", zipDownloadProgress)
//			}
//		case gameDownloadProgress, ok := <-gameDownloadChan:
//
//			if !ok {
//				fmt.Println("Game Download Complete 👍")
//				gameDownloadChan = nil
//			} else {
//				fmt.Printf("game Download Progress := %d \n", gameDownloadProgress)
//			}
//
//		case documentDownloadProgress, ok := <-documentDownloadChan:
//
//			if !ok {
//				fmt.Println("Document Download Complete 👍 ")
//				documentDownloadChan = nil
//			} else {
//				fmt.Printf("documet Download Progress := %d \n", documentDownloadProgress)
//			}
//
//		case <-time.After(100 * time.Second):
//			fmt.Println("Download Speed Very slow exitting ")
//			return
//
//		}
//
//		if gameDownloadChan == nil && movieDownlodChan == nil && zipDownloadChan == nil && documentDownloadChan == nil {
//			fmt.Println("All files Downloaded Successfully, !! Exiting the loop")
//			break
//		}
//	}
//
//	wg.Wait()
//}
