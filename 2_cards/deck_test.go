package main

import (
	"testing"
	"os"
)

func Test_new_deck(t *testing.T) {
	d := new_deck()
	length := len(d)
	first_card := d[0]
	last_card := d[length-1]

	if length != 16 {
		t.Errorf("Expected deck length of 16, but got %v", length)
	}

	if first_card != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", first_card)
	}

	if last_card != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", last_card)
	}
}

func Test_save_to_file_and_new_deck_from_file(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := new_deck()
	d.save_to_file(filename)

	loaded_deck := new_deck_from_file(filename)

	if len(loaded_deck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loaded_deck))
	}

	os.Remove(filename)
}