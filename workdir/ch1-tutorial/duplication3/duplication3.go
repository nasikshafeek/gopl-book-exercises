package main

import (
	"fmt"
	"os"
	"strings"
)

/**
Main difference between this file and the duplication2 is that,
this logic of reading the file is it reads the file in it's entirety and loads it completely into memory.
Whereas the duplication2 logic is it reads in a line by line under the "streaming" mode.
*/

func main() {
	counts := make(map[string] int)
	fileForLine := make(map[string] []string)
	for _, fileName := range os.Args[1:] {
		// Read the file
		readFileData, fileReadErr := os.ReadFile(fileName)
		if fileReadErr != nil {
			fmt.Fprint(os.Stderr, "Error reading file:", fileName, fileReadErr.Error())
		}

		// Use the file byte data to read line by line
		for _, line := range strings.Split(string(readFileData), "\n") {
			if line == "" {
				continue
			}
			counts[line]++
			fileForLine[line] = append(fileForLine[line], fileName)
		}
	}

	// print the counts infromation
	for line, count := range counts {
		var fileNames = strings.Join(fileForLine[line], ", ") // Would contain duplicates
		fmt.Printf("%d\t%s\t%s\n", count, line, fileNames)
	}
}