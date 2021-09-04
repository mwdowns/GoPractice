package main

import (
	"fmt"

	"github.com/mwdowns/GoPractice/modern_webapp/my_packages/helpers"
)

func main() {
	var housePets []helpers.Pet
	housePets = append(housePets, helpers.Pet{Type: "dog", Name: "Danny", Word: "Buscuits?", Legs: 4})
	housePets = append(housePets, helpers.Pet{Type: "dog", Name: "Rovi", Word: "Huh?", Legs: 4})
	housePets = append(housePets, helpers.Pet{Type: "cat", Name: "Stinkins", Word: "Lovins?", Legs: 4})
	housePets = append(housePets, helpers.Pet{Type: "cat", Name: "Doodins", Word: "Fancy Feast?", Legs: 4})
	for _, pet := range housePets {
		PrintPetInfo(&pet)
	}
}

func PrintPetInfo(a helpers.Animal) {
	fmt.Println(a.SayName(), "says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}
