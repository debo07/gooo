package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// This is a custom type
type deck []string

// Create a new function that returns a deck
func newDec() deck {
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

func (d deck) print() {
	for i, card := range d {
		fmt.Println( i, card )
	}
}