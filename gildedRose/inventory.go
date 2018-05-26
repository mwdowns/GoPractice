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
	EpicItems    []item
	SpecialItems []item
	IonizedItems []item
	NormalItems  []item
}

func newItem(name string, sellby int, quality int) item {
	return item{Name: name, SellBy: sellby, Quality: quality}
}

func (i item) printItem() {
	fmt.Printf("Name: %v\n", i.Name)
	fmt.Printf("Sell in %v days.\n", i.SellBy)
	fmt.Printf("Quality: %v\n", i.Quality)
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
	if len(i.IonizedItems) == 0 {
		fmt.Println("There are no Ionized Items. Try back later.")
	} else {
		fmt.Printf("Ionized Items: %v\n", i.IonizedItems)
	}
	if len(i.NormalItems) == 0 {
		fmt.Println("There are no Normal Items. Try back later.")
	} else {
		fmt.Printf("Normal Items: %v\n", i.NormalItems)
	}
}

func (i inventory) countInventory() {
	ic := len(i.EpicItems) + len(i.SpecialItems) + len(i.IonizedItems) + len(i.NormalItems)
	fmt.Printf("There are %v items in the shop!\n", ic)
}

func inventoryRandomizer(num int) inventory {
	randomInventory := inventory{}
	Objects := []string{"Blaster", "Med-Pack", "Nano-Injector", "Force-sheild", "Gloves", "Boots"}
	Adjectives := []string{"Awesome", "Good", "Excellent", "Bodacious", "OK", "Fine"}
	Attributes := []string{"Velocity", "Energy", "Kismet", "Recovery", "Cloaking", "Titanium", "Plasma"}
	Cheeses := []string{"Brie", "Cheddar", "Gouda"}
	for i := 0; i < num; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		object := Objects[r.Intn(len(Objects)-1)]
		r = rand.New(source)
		adj := Adjectives[r.Intn(len(Adjectives)-1)]
		r = rand.New(source)
		attr := Attributes[r.Intn(len(Attributes)-1)]
		r = rand.New(source)
		isCheese := r.Intn(3)
		r = rand.New(source)
		cheese := Cheeses[r.Intn(len(Cheeses)-1)]
		cheeseName := "Aged " + cheese + " of " + attr
		itemName := object + " of " + adj + " " + attr
		r = rand.New(source)
		itemType := r.Intn(100)
		if itemType == 1 {
			itemName = object + " of Ultimate " + attr
			randomInventory.EpicItems = append(randomInventory.EpicItems, item{Name: itemName, SellBy: 10, Quality: 80})
		} else if itemType < 5 {
			if isCheese == 2 {
				randomInventory.SpecialItems = append(randomInventory.SpecialItems, item{Name: cheeseName, SellBy: 10, Quality: 5})
			} else {
				itemName = itemName + " and Increasing Quality"
				randomInventory.SpecialItems = append(randomInventory.SpecialItems, item{Name: itemName, SellBy: 10, Quality: 5})
			}
		} else if itemType <= 20 {
			itemName = "Ionized " + itemName
			randomInventory.IonizedItems = append(randomInventory.IonizedItems, item{Name: itemName, SellBy: 10, Quality: 20})
		} else {
			itemName = object + " of Somewhat " + adj + " " + attr
			randomInventory.NormalItems = append(randomInventory.NormalItems, item{Name: itemName, SellBy: 10, Quality: 10})
		}
	}
	return randomInventory
}
