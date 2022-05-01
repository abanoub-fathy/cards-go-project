package main

func main() {
	// create new deck
	deck1 := CreateNewDeck()
	// shuffle the deck
	deck1.Shuffle()
	// print deck
	deck1.PrintAllCards()
}
