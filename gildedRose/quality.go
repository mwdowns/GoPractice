package main

import (
	"strings"
)

func (i *item) decreaseQuality(isConjured bool) {
	if isConjured == true {
		i.Quality = i.Quality - 2

	} else {
		i.Quality = i.Quality - 1
	}
}

func (i *item) increaseQuality() {
	i.Quality = i.Quality + 1
	if strings.Contains(i.Name, "Aged") == false {
		if i.SellBy < 6 {
			i.Quality = i.Quality + 1
		}
		if i.SellBy < 3 {
			i.Quality = i.Quality + 3
		}
		if i.SellBy == 0 {
			i.Quality = 0
		}
	}
}
