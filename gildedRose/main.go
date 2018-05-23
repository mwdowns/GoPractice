package main

import (
	"fmt"
)

func main() {

	fmt.Println("Doin' the guilded rose thing in Go!")
	fmt.Println("Want to play the game? (y/n) ")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "y" {
		fmt.Println("Great")
		startGame()
	} else if answer == "n" {
		fmt.Println("OK. See you later")
	} else {
		fmt.Println("Please choose y or n")
		main()
	}
}

func startGame() {
	inventory := inventoryRandomizer(15)
	inventory.printInventory()
}
