package main

import (
	"os"
	"testing"
)

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

func TestSaveTofileAndGetDeckFromFile(t *testing.T) {
	os.Remove("_testDeckFile")

	// Created and saved a file
	deck := newDeck()
	err := deck.saveTofile("_testDeckFile")

	if err != nil {
		t.Errorf("Found Error: %v while saving the deck to file, expected no error", err)
	}

	// Read from file
	loadedDeck := getDeckFromFile("_testDeckFile")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in the loaded deck, received %v", len(loadedDeck))
	}

	// Finally cleanup: delete the file created while testing
	os.Remove("_testDeckFile")
}