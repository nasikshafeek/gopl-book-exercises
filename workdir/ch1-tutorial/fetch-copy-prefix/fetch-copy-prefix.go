package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
 * Adding a conditional to fetch-copy to find out if there exists the protocol in the URL, otherwise add it.
 */
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch-copy-prefix: %v\n", err)
			os.Exit(1)
		}

		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "fetch-copy-prefix: reading %s: %v\n", url, copyErr)
			os.Exit(1)
		}

		fmt.Printf("\nStatus code for the response: %v\n", resp.StatusCode)
		fmt.Println("\n\n\n\n\n")
	}
}