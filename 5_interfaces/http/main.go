package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type log_writer struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit((1))
	}

	// bs := make([]byte, 99999) // Create a byte slice with 99999 elements
	// resp.Body.Read(bs) // Read the response body into the byte slice
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	lw := log_writer{}
	io.Copy(lw, resp.Body)

}

func (log_writer) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many", len(bs), "bytes.")
	return len(bs), nil
}