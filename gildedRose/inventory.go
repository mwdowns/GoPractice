package main

import (
	"fmt"
	"math/rand"
	"time"
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
	if len(i.EpicItems) == 0 {
		fmt.Println("There are no Epic Items. Try back later.")
	} else {
		fmt.Printf("Epic Items: %v\n", i.EpicItems)
	}
	if len(i.SpecialItems) == 0 {
		fmt.Println("There are no Special Items. Try back later.")
	} else {
		fmt.Printf("Special Items: %v\n", i.SpecialItems)
	}
	if len(i.ConjuredItems) == 0 {
		fmt.Println("There are no Conjured Items. Try back later.")
	} else {
		fmt.Printf("Conjured Items: %v\n", i.ConjuredItems)
	}
	if len(i.NormalItems) == 0 {
		fmt.Println("There are no Normal Items. Try back later.")
	} else {
		fmt.Printf("Normal Items: %v\n", i.NormalItems)
	}
}

func (i item) printItem() {
	fmt.Printf("Name: %v\n", i.Name)
	fmt.Printf("Sell in %v days.\n", i.SellBy)
	fmt.Printf("Quality: %v\n", i.Quality)
}

func inventoryRandomizer(seed int) inventory {
	randomInventory := inventory{}
	Objects := []string{"Sword", "Potion", "Salve", "Sheild", "Gauntlets", "Boots"}
	Adjectives := []string{"Awesome", "Good", "Excellent", "Bodacious", "OK", "Fine"}
	Attributes := []string{"Speed", "Power", "Luck", "Healing", "Invisibility", "Stone", "Fire"}
	for i := 0; i < seed; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		object := Objects[r.Intn(len(Objects)-1)]
		adj := Adjectives[r.Intn(len(Adjectives)-1)]
		attr := Attributes[r.Intn(len(Attributes)-1)]
		itemName := object + " of " + adj + " " + attr
		itemType := r.Intn(100)
		if itemType == 1 {
			itemName = object + "of Ultimate " + attr
			randomInventory.EpicItems = append(randomInventory.EpicItems, item{Name: itemName, SellBy: 10, Quality: 80})
		}
		if itemType < 5 {
			itemName = itemName + " of Increasing Quality"
			randomInventory.SpecialItems = append(randomInventory.SpecialItems, item{Name: itemName, SellBy: 10, Quality: 5})
		}
		if itemType < 20 {
			itemName = "Conjured " + itemName
			randomInventory.ConjuredItems = append(randomInventory.ConjuredItems, item{Name: itemName, SellBy: 10, Quality: 20})
		} else {
			itemName = object + " of Somewhat " + adj + " " + attr
			randomInventory.NormalItems = append(randomInventory.NormalItems, item{Name: itemName, SellBy: 10, Quality: 10})
		}
	}
	return randomInventory
}
