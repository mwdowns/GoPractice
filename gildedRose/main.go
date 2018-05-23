package main

import "fmt"

func main() {

	swordOfPrettyGoodStats := newItem("Sword of Pretty Good Stats", 10, 10)

	// inventory := inventoryRandomizer(15)

	fmt.Println("Doin' the guilded rose thing in Go!")
	// inventory.countInventory()
	// inventory.printInventory()

	swordOfPrettyGoodStats.printItem()
	swordOfPrettyGoodStats.anotherDayCloserToDeath()
	swordOfPrettyGoodStats.printItem()
}
