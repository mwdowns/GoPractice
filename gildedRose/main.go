package main

import "fmt"

func main() {

	swordOfPrettyGoodStats := newItem("Sword of Pretty Good Stats", 10, 10)

	inventory := inventoryRandomizer(10)
	inventory2 := newInventory()
	inventory.NormalItems = append(inventory.NormalItems, swordOfPrettyGoodStats)

	fmt.Println("Doin' the guilded rose thing in Go!")
	inventory2.printInventory()
	swordOfPrettyGoodStats.printItem()
	inventory.printInventory()
}
