package main

func main() {
	// Create a new deck
	cardList := newDeck()

	hand, remainingDeck := deal(cardList, 5)

	hand.print()
	remainingDeck.print()
}
