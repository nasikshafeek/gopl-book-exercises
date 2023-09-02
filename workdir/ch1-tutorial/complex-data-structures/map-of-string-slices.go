package main

import "fmt"

func main() {
    x := make(map[string][]string)

    x["key"] = append(x["key"], "value")
    x["key"] = append(x["key"], "value1")

    fmt.Println(x["key"][0])
    fmt.Println(x["key"][1])
}