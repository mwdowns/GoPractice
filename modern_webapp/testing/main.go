package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := divideThings(100.00, 0.00)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("the answer is", r)
}

func divideThings(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("division by 0")
	}

	result = x / y
	return result, nil
}
