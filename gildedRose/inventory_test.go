package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewItem(t *testing.T) {
	i := newItem("New Item", 10, 10)
	if reflect.TypeOf(i) != reflect.TypeOf(item{}) {
		t.Errorf("Expected i to be a item of type struct, got %v", reflect.TypeOf(i))
	}
	fmt.Println("Test for newItem function passed")
}

func TestNewInventory(t *testing.T) {
	i := newInventory()
	if reflect.TypeOf(i) != reflect.TypeOf(inventory{}) {
		t.Errorf("Expected i to be of inventory type struct, got %v", reflect.TypeOf(i))
	}
	fmt.Println("Test for newInventory function passed")
}
