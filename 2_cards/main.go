package main

// Valid
// var num int
//  Invalid
// num2 := 1

func main() {
	// var card string = "Ace of Spades"
	cards := new_deck()
	cards.shuffle()
	cards.print()
}