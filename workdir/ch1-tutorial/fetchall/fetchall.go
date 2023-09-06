package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
		go fetch(url, channel1) // start a goroutine to fetch the URL
	}

	// The following listens to the messenger channel for all the data received and then prints them
	for range os.Args[1:] {
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

	nBytes, copyErr := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if copyErr != nil {
		channel <- fmt.Sprint(os.Stderr, "fetch-copy-prefix: reading %s: %v\n", url, copyErr)
		return
	}
	
	secondsElapsed := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", secondsElapsed, nBytes, url)
}