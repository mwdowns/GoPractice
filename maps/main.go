package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "00ff00",
	}

	colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {
	for _, hex := range c {
		fmt.Println(hex)
	}
}
