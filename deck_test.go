package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d1 := CreateNewDeck()

	// test the length
	if len(d1) != 16 {
		t.Errorf("Expected 16 but got %v", len(d1))
	}

	// test the first card
	if d1[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", d1[0])
	}

	// test last element
	if d1[len(d1)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs but got %v", d1[len(d1)-1])
	}

}

func TestCreateAndReadDeckFromFile(t *testing.T) {
	fileName := "_deckTesting.txt"
	// remove the file
	os.Remove(fileName)

	// create new deck and save it to a file
	d1 := CreateNewDeck()
	d1.SaveToFile(fileName)

	// create second deck from that file
	d2 := CreateDeckFromFile(fileName)

	// remove the file
	os.Remove(fileName)

	if len(d2) != len(d1) {
		t.Errorf("Expected len=%v but got len=%v", len(d1), len(d2))
	}
}
