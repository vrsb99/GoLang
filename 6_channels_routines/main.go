package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, line := range links {
		go checkLink(line, c)
	}

	// for {go checkLink(<-c, c)}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}