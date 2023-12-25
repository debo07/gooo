package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
// This is a custom type
type deck []string

// Create a new function that returns a deck
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suite)
		}
	}
	return cards
}

// Create a new function that prints a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println( i, card )
	}
}

// Create a new function that deals a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save the deck to a file
func (d deck) saveTofile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// retrieve a deck fro file
func getDeckFromFile(fileName string) (d deck) {
	bs, err := os.ReadFile(fileName)
	if(err != nil) {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Convert byte slice to deck
	// byte slice -> string -> string slice -> deck
	d = deck(strings.Split(string(bs), ","))
	return
}

// shuffle an existing deck
func (d deck) shuffle() {
	for i := range(d) {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}