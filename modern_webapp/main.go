package main

import (
	"log"
)

type myVar struct {
	Blah string
}

func main() {
	myMap := make(map[string]myVar)
	mySlice := []myVar{}


	blah := myVar{Blah: "blah",}
	bla2 := myVar{Blah: "blah2",}

	myMap["dog"] = blah
	mySlice = append(mySlice, blah)
	mySlice = append(mySlice, bla2)

	log.Println(myMap["dog"].Blah)
	log.Println(mySlice[0].Blah, mySlice[1].Blah)
}
