package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run workdir/ch1-tutorial/duplication2/duplication2.go ./abc ./xyz

func main() {
	duplicateCounts := make(map[string]int)
	files := os.Args[1:]

	if (len(files) == 0) {
		countLines(os.Stdin, duplicateCounts)
	} else {
		for _, filepath := range files {
			aFile, openErr := os.Open(filepath)

			// Test this error case by revoking read permissions from one of the text file paths passed in as arguments.
			if openErr != nil {
				fmt.Fprintf(os.Stderr, "error at duplication2: %v\n", openErr)
				continue
			}

			countLines(aFile, duplicateCounts)
			aFile.Close()
		}
	}

	printResults(duplicateCounts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		/**
		In the world of maps if a key is non existend it will be created by running the following code
		Usually when we do ++ operation it will do it on the zero value for the given type (for int it's 0)
		*/
		counts[input.Text()]++ // How would an empty value be incremented/handled?
	}
}

func printResults(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}