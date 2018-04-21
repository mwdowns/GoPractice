package main

import (
	"fmt"
)

// create a new type of deck
// which is a slice of strings

type card struct {
	name  string
	suit  string
	value string
}

type deck []card

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{value + " of " + suit, suit, value})
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) cardPrint() {
	for index, card := range d {
		fmt.Println(index, card.name)
	}
}

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

// func (d deck) saveToFile(fileName string) error {
// 	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
// }

// func newDeckFromFile(fileName string) deck {
// 	bs, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println(":::Error::: ", err)
// 		os.Exit(1)
// 	}
// 	s := strings.Split(string(bs), ",")
// 	return deck(s)
// }

// func (d deck) shuffle() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)
// 	for i := range d {
// 		//generate a random number from -1 to len of deck
// 		newPosition := r.Intn(len(d) - 1)
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }
