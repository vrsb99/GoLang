package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, text)
}