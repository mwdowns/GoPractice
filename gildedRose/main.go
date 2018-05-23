package main

import (
	"fmt"
)

var answer string

func main() {

	fmt.Println("Doin' the guilded rose thing in Go!")
	fmt.Println("Want to play the game? (y/n) ")
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
	fmt.Println("Would you like to create a player? (y/n)")
	fmt.Scanf("%s", &answer)
	if answer == "y" {
		fmt.Println("Great")
		p := createNewPlayer()
		p.printPlayer()
	} else if answer == "n" {
		fmt.Println("OK. See you later")
	} else {
		fmt.Println("Please choose y or n")
		startGame()
	}
}
