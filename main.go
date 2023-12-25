package main

import "fmt"

func main() {
	// Create a new deck
	cardList := newDeck()
	fmt.Println("----------- Original Deck -----------")
	cardList.print()

	hand, remainingDeck := deal(cardList, 5)

	fmt.Println("----------- Hand -----------")
	hand.print()
	fmt.Println("----------- Remaining Deck -----------")
	remainingDeck.print()

	// Write the newly created deck to a file.
	hand.saveTofile("hand.txt")
	remainingDeck.saveTofile("remaining_deck.txt")

	// Read a new deck from file
	deckFromFile := getDeckFromFile("hand.txt")
	fmt.Println("----------- Deck from file -----------")
	deckFromFile.print()

	// Shuffle a deck and print
	cardList.shuffle()
	fmt.Println("----------- Deck after Shuffling -----------")
	cardList.print()
}
