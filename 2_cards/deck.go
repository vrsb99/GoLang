package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func new_deck() deck {
	cards := deck{}
	card_suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	card_values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range card_suits {
		for _, value := range card_values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// By convention, the receiver variable is the first letter of the type
// Extension method
func (d deck) print() {
	for i, per_card := range d {
		fmt.Println(i, per_card)
	}
}

// Argument method
func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

func (d deck) to_string() string {
	return strings.Join([]string(d), ",")
}

func (d deck) save_to_file(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.to_string()), 0666)
}

func new_deck_from_file(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// For slices, pointers are not required as they are already references of the underlying array
func (d deck) shuffle() {
	var length int = len(d) - 1
	var source rand.Source = rand.NewSource(time.Now().UnixMilli())
	var r *rand.Rand = rand.New(source)
	
	for index := range d {
		new_position := r.Intn(length) 
		d[index], d[new_position] = d[new_position], d[index]
	}
}