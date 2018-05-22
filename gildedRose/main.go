package main

import "fmt"

func main() {

	swordOfPrettyGoodStats := newItem("Sword of Pretty Good Stats", 10, 10)

	inventory := newInventory()

	inventory.NormalItems = append(inventory.NormalItems, swordOfPrettyGoodStats)

	fmt.Println("Doin' the guilded rose thing in Go!")
	inventory.printInventory()
}
