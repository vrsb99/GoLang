package main

import (
	"fmt"
)

type bot interface {
	get_greeting() string
	//get_greeting(string, int) (string, error)
}

type english_bot struct{}
type spanish_bot struct{}


func main() {
	eb := english_bot{}
	sb := spanish_bot{}

	print_greeting(eb)
	print_greeting(sb)
}

func (eb english_bot) get_greeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanish_bot) get_greeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}

func print_greeting(b bot) {
	fmt.Println(b.get_greeting())
}