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
		// p := createNewPlayer()
		p := createNewPlayerDialogue()
		p.printPlayer()
	} else if answer == "n" {
		fmt.Println("OK. See you later")
	} else {
		fmt.Println("Please choose y or n")
		fmt.Println("Would you like to create a player? (y/n)")
		startGame()
	}
}

func createNewPlayerDialogue() *player {
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
	return createNewPlayer(name, class)
}
