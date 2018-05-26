package main

import (
	"fmt"
)

type player struct {
	Name      string
	Class     string
	Stats     stats
	Inventory inventory
	Weapon    item
	Armor     item
}

type stats struct {
	Itl int
	Chr int
	Mgt int
	Dex int
}

func createNewPlayer() *player {
	var name string
	var class string
	fmt.Println("What name would you like to give your player? ")
	fmt.Scanln("%s", &name)
	fmt.Println("What class would you like to give your player? (for a list of classes, type 'help')")
	fmt.Scanln("%s", &class)
	if class == "help" {
		fmt.Println("Scoundrel - Engineer - Captain")
		fmt.Println("Which class would you like to play? ")
		fmt.Scanln("%s", &class)
	}
	fmt.Println("Great! You're ready to start your adventure!")
	p := &player{
		Name:  name,
		Class: class,
		Stats: stats{
			Itl: 10,
			Chr: 10,
			Mgt: 10,
			Dex: 10,
		},
		Inventory: inventory{},
		Weapon:    item{},
		Armor:     item{},
	}
	p.createPlayerInventory()
	p.Weapon = p.Inventory.NormalItems[1]
	p.Armor = p.Inventory.NormalItems[2]
	return p
}

func (p *player) printPlayer() {
	fmt.Printf("Your player's name is %v.\n", p.Name)
	fmt.Printf("They are a %v class.\n", p.Class)
	fmt.Printf("Their maing weapon is a %v.\n", p.Weapon.Name)
	fmt.Printf("They are wearing a %v.\n", p.Armor.Name)
	fmt.Println("Their inventory consists of these items:")
	for _, i := range p.Inventory.NormalItems {
		fmt.Println(i.Name)
	}
	fmt.Printf("They have on special item %v", p.Inventory.SpecialItems[0].Name)
}

func (p *player) createPlayerInventory() {
	food := newItem("Space Food of Speedy Recovery", 20, 10)
	blaster := newItem("Plasma Blaster of Pretty Good Stats", 1000, 10)
	hat := newItem("Fighter Helmet of Thick Headedness", 1000, 10)
	cheese := newItem("Good Aged Gouda", 10, 5)
	p.Inventory = inventory{
		EpicItems: []item{},
		SpecialItems: []item{
			cheese,
		},
		IonizedItems: []item{},
		NormalItems: []item{
			food,
			blaster,
			hat,
		},
	}
}
