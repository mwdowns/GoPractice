package main

import (
	"fmt"
)

type item struct {
	Name    string
	SellBy  int
	Quality int
}

type inventory struct {
	EpicItems     []item
	SpecialItems  []item
	ConjuredItems []item
	NormalItems   []item
}

func newItem(name string, sellby int, quality int) item {
	return item{Name: name, SellBy: sellby, Quality: quality}
}

func newInventory() inventory {
	return inventory{}
}

func (i inventory) printInventory() {
	fmt.Println(i)
}
