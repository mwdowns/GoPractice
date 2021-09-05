package main

import (
	"encoding/json"
	"fmt"
	"modern_webapp/refreshers/helpers"
)

const Num = 10

func GetRandomNumber(numChan chan int) {
	randomNum := helpers.GenerateRandomNumber(Num)
	numChan <- randomNum
}

func main() {
	housePets := []helpers.Pet{}
	response := `
		[
			{
				"type": "dog",
				"name": "Danny",
				"word": "Buscuits?",
				"legs": 4
			},
			{
				"type": "dog",
				"name": "Rovi",
				"word": "Huh?",
				"legs": 4
			},
			{
				"type": "cat",
				"name": "Stinkins",
				"word": "Lovins?",
				"legs": 4
			},
			{
				"type": "cat",
				"name": "Doodins",
				"word": "Fancy Feast?",
				"legs": 4
			}
		]
	`
	err := json.Unmarshal([]byte(response), &housePets)
	if err != nil {
		fmt.Println("got error")
	}
	// housePets = append(housePets, helpers.Pet{Type: "dog", Name: "Danny", Word: "Buscuits?", Legs: 4})
	// housePets = append(housePets, helpers.Pet{Type: "dog", Name: "Rovi", Word: "Huh?", Legs: 4})
	// housePets = append(housePets, helpers.Pet{Type: "cat", Name: "Stinkins", Word: "Lovins?", Legs: 4})
	// housePets = append(housePets, helpers.Pet{Type: "cat", Name: "Doodins", Word: "Fancy Feast?", Legs: 4})
	for _, pet := range housePets {
		PrintPetInfo(&pet)
	}

	newJson, err := json.MarshalIndent(housePets, "", "  ")
	if err == nil {
		fmt.Println(string(newJson))
	}

	numChan := make(chan int)
	defer close(numChan)

	go GetRandomNumber(numChan)
	num := <-numChan
	fmt.Println(num)
}

func PrintPetInfo(a helpers.Animal) {
	fmt.Println(a.SayName(), "says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}
