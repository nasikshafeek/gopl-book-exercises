package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run workdir/ch1-tutorial/duplication2/duplication2.go ./abc ./xyz

func main() {
	duplicateCounts := make(map[string]int)
	duplicatedFileNames := make(map[string] []string)
	files := os.Args[1:]

	if (len(files) == 0) {
		countLines(os.Stdin, duplicateCounts, duplicatedFileNames)
	} else {
		for _, filepath := range files {
			aFile, openErr := os.Open(filepath)

			// Test this error case by revoking read permissions from one of the text file paths passed in as arguments.
			if openErr != nil {
				fmt.Fprintf(os.Stderr, "error at duplication2: %v\n", openErr)
				continue
			}

			countLines(aFile, duplicateCounts, duplicatedFileNames)
			aFile.Close()
		}
	}

	printResults(duplicateCounts, duplicatedFileNames)
}

func countLines(f *os.File, counts map[string]int, duplicatedFileNames map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		/**
		In the world of maps if a key is non existend it will be created by running the following code
		Usually when we do ++ operation it will do it on the zero value for the given type (for int it's 0)
		*/
		counts[input.Text()]++ // How would an empty value be incremented/handled?
		
		duplicatedFileNames[input.Text()] = append(duplicatedFileNames[input.Text()], f.Name())
	}
}

func printResults(counts map[string]int,  duplicatedFileNames map[string][]string) {
	for line, count := range counts {
		if count > 1 {
			// Deduplicating the file names if there are multiple files for a given line.
			var uniqueFiles []string
			uniqueFileNames :=  make(map[string] bool, 0)
			for _, fileName := range duplicatedFileNames[line] {
				uniqueFileNames[fileName] = true
			}
			for fileName := range uniqueFileNames {
				uniqueFiles = append(uniqueFiles, fileName)
			}
			
			fmt.Printf("%d\t%s\t%s\n", count, line, strings.Join(uniqueFiles, ", "))
		}
	}
}