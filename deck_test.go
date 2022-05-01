package main

import "testing"

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
