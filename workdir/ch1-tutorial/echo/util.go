package main

import "fmt"

func printEcho(parameters ...any) {
	if len(parameters) == 1 {
		fmt.Println(parameters[0])
	} else {
		fmt.Print(parameters[1], ": ", parameters[0], "\n")
	}
}