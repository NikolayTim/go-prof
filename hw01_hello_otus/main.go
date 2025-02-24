package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	str := reverse.String("Hello, OTUS!")

	fmt.Println(str)
}
