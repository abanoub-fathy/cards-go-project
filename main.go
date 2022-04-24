package main

import "fmt"

func main() {
	// create new deck
	cards1 := CreateNewDeck()
	// print the deck
	fmt.Println(cards1)
	// save the deck onto file
	cards1.SaveToFile("cards1.txt")

	// create new deck from a file
	cards2 := CreateDeckFromFile("cards1.txt")
	// print the deck
	fmt.Println(cards2)

}
