package main

import "fmt"

func main() {

	// swordOfPrettyGoodStats := newItem("Sword of Pretty Good Stats", 10, 10)

	inventory := inventoryRandomizer(15)
	// inventory2 := newInventory()
	// inventory2.NormalItems = append(inventory.NormalItems, swordOfPrettyGoodStats)

	fmt.Println("Doin' the guilded rose thing in Go!")
	// inventory2.printInventory()
	// swordOfPrettyGoodStats.printItem()
	inventory.countInventory()
	// inventory.printInventory()
}
