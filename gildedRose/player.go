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

func createNewPlayer() player {
	var name string
	var class string
	fmt.Println("What name would you like to give your player? ")
	fmt.Scanln("%s", &name)
	fmt.Println("What class would you like to give your player? (for a list of classes, type 'help'")
	fmt.Scanln("%s", &class)
	if class == "help" {
		fmt.Println("Scoundrel - Engineer - Captain")
		fmt.Println("Which class would you like to play? ")
		fmt.Scanln("%s", &class)
	}
	fmt.Println("Great! You're ready to start your adventure!")
	p := player{
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
