package main

import "fmt"

// create a new type of deck
// which is a slice of strings

type deck []string

// type card struct {
// 	name string
// 	value int
// }

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) cardPrint() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
