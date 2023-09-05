package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
 * Fetch but non blocking as Copy uses streaming than waiting to read the complete response.
 */
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch-copy: %v\n", err)
			os.Exit(1)
		}
		
		// bytes, err := io.ReadAll(resp.Body)
		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "fetch-copy: reading %s: %v\n", url, copyErr)
			os.Exit(1)
		}
	}
}