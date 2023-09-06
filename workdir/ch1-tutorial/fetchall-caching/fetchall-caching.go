package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
 * Fetchall fetches URLs in parallel and report their times and sizes
 */
func main() {
	startTime := time.Now()
	channel1 := make(chan string)

	for _, url := range os.Args[1:] {
		// asynchronous call to fetch the URL along with the messenger channel
		go fetch(url, channel1)
		go fetch(url, channel1)
	}

	for i := 0; i < 2 * len(os.Args[1:]); i ++{
		fmt.Println(<-channel1) // recieve from the channel channel1
	}

	fmt.Printf("%.2fs of total elapsed time\n", time.Since(startTime).Seconds())
}

func fetch(url string, channel chan<- string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err) // Send message to the channel
		return
	}

	var fileName string

	if strings.HasPrefix(url, "https://") {
		fileName = strings.SplitAfter(url, "https://")[1]
	} else if strings.HasPrefix(url, "http://") {
		fileName = strings.SplitAfter(url, "http://")[1]
    }
	
	fileName += strconv.Itoa(time.Now().Nanosecond())
	fileName += ".html"

	fileDestination, fileCreationError := os.Create("./outputs/" + fileName)
	if fileCreationError != nil {
		channel <- fmt.Sprintf("fetchall-caching error: %s %v\n", fileName, fileCreationError)
		return
	}

	nBytes, copyErr := io.Copy(fileDestination, resp.Body)
	resp.Body.Close()

	if copyErr != nil {
		channel <- fmt.Sprintf("fetch-copy-prefix: reading %s %v\n", url, copyErr)
		return
	}
	
	secondsElapsed := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", secondsElapsed, nBytes, url)
}