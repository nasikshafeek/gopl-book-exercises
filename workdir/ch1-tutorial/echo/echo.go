package main

import (
	"os"
	"strconv"

	. "gopl.io/workdir/ch1-tutorial/custom-package"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println("Length:", len(os.Args[1:]))
	// loopThroughToReconstructArgument()
	// forEachLoopThroughToReconstructArgument()
	SayHello("Nasik")
	printArgumentIndexAndValue()
}

// Loop over the arguments to re-construct the echo output
func loopThroughToReconstructArgument() {
	var str, separator string
	for i := 1; i < len(os.Args); i++ {
		str += separator + os.Args[i]
		separator = " "
	}
	printEcho(str, os.Args[0])
}

func printArgumentIndexAndValue() {
	str, separator := "", " "
	for arg_index, arg := range os.Args[1:] {
		str += arg + separator
		separator = " "
		printEcho(arg, "Argument " + strconv.Itoa(arg_index + 1))
	}
}
