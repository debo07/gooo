package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected the deck length of 16, but got %v", len(d))
	}

	// Add checks for 1st and last card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades', but received %v", d[0])
	}
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs', but received %v", d[len(d) - 1])
	}
}