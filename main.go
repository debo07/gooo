package main

func main() {
	// Create a new deck
	cardList := newDeck()

	hand, remainingDeck := deal(cardList, 5)

	hand.print()
	remainingDeck.print()

	// Write the newly created deck to a file.
	hand.saveTofile("hand.txt")
	remainingDeck.saveTofile("remaining_deck.txt")
}
