package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// define a deck type which is []string
type deck []string

// method to create a new deck
func CreateNewDeck() deck {
	cards := deck{}

	// suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}

// method to print all the cards in the deck
func (d deck) PrintAllCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// method to deal
// This method returns the hands cards and the remaining cards
func (d deck) Deal(handNum int) (deck, deck) {
	return d[:handNum], d[handNum:]
}

// method to convert a deck to string
func (d deck) ConvertToString() string {
	return strings.Join(d, ",")
}

// method to save Deck to a file
func (d deck) SaveToFile(fileName string) {
	err := ioutil.WriteFile(fileName, []byte(d.ConvertToString()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// method to create a new deck from file
func CreateDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// convert the bs to string
	content := string(bs)

	// convert content to []string
	cards := strings.Split(content, ",")

	// return deck
	return deck(cards)
}

// method to shuffle cards
func (d deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	// loop through all elemnts in the deck
	for i := range d {
		// generate randIndex
		randIndex := rand.Intn(len(d) - 1)
		fmt.Printf("i = %d, rand=%d\n", i, randIndex)
		// swap two elements in the array
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}
