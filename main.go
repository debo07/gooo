package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	// Slice - Defining a slice of type string
	cardList := []string{"Ace of Diamonds", newCard()}
	// Appending a new element to the slice
	cardList = append(cardList, "Six of Spades")

	// Iterate over the slice
	for i, card := range cardList {
		fmt.Println(i, card)
	}
}
