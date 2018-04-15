package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remaingDeck := deal(cards, 5)
	hand.cardPrint()
	fmt.Printf("remaining deck\n")
	remaingDeck.cardPrint()
}
