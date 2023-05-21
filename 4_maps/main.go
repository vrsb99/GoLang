package main

// All keys must be of the same type
// All values must be of the same type

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	// fmt.Println(colors)

	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	print_map(colors)
}

func print_map(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}