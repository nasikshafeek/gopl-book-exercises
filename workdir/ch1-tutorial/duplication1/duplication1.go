package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		/**
		In the world of maps if a key is non existend it will be created by running the following code
		Usually when we do ++ operation it will do it on the zero value for the given type (for int it's 0)
		*/
		counts[input.Text()]++ // How would an empty value be incremented/handled?
	}

	for line, count := range counts {
		// if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		// }
	}
}