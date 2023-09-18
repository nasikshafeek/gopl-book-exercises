package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/**
 * Fetchall fetches URLs in parallel and report their times and sizes
 */
func main() {
	// startTime := time.Now()
	// fetchChannel := make(chan string)
	csvReaderChannel := make(chan string)

	go readCsv("./top-1m.csv", csvReaderChannel)

	// Ranging over a channel is an non terminal loop, until it receives a close message.
	for channelMsg := range csvReaderChannel { 
		fmt.Println("url from csvReaderChannel", channelMsg)
	}

	// for _, url := range os.Args[1:] {
	// 	// asynchronous call to fetch the URL along with the messenger channel
	// 	go fetch(url, fetchChannel) // start a goroutine to fetch the URL
	// }

	// // The following listens to the messenger channel for all the data received and then prints them
	// for range os.Args[1:] {
	// 	fmt.Println(<-fetchChannel) // recieve from the channel channel1
	// }

	// fmt.Printf("%.2fs of total elapsed time\n", time.Since(startTime).Seconds())
}

/*
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
*/

func readCsv(filename string, readerChannel chan<- string) {
	csvFile, err := os.Open(filename)

    if err!= nil {
        fmt.Printf("readCsv: %v\n", err)
        return
    }
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	if _, readingHeadErr := csvReader.Read(); readingHeadErr != nil {
		fmt.Fprintln(os.Stderr, "reading-head-error: ", readingHeadErr)
		return
	}

	// chunk, chunkReadError := csvReader.Read()
	// if chunkReadError != nil {
	// 	fmt.Fprintln(os.Stderr, "reading-chunk-error: ", chunkReadError)
	// }

	// fmt.Println("Read chunk", url, chunkReadError);

	// For each row in the CSV file, read the url and send it to the channel,
	// This channel message should be gathered by the main function and be used to fetch the URL.
	// record, readErr := csvReader.Read()
	for record, readErr := csvReader.Read(); readErr != io.EOF; record, readErr = csvReader.Read() {
		url := record[1]
		readerChannel <- url

		// record, readErr = csvReader.Read()
		// chunk, chunkReadError := csvReader.Read()
		// fmt.Println("Read chunk", url, readErr);

		// chunk, chunkReadError = csvReader.Read()
		// fmt.Println("Read chunk", chunk, chunkReadError);
	}

	// Always close the channel from the sender, not the receiver
	close(readerChannel)
}