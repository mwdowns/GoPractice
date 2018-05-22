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

func TestInventoryRandomizer(t *testing.T) {
	i := inventoryRandomizer(15)
	itemCount := len(i.EpicItems) + len(i.SpecialItems) + len(i.ConjuredItems) + len(i.NormalItems)
	if itemCount != 15 {
		t.Errorf("Expected 15 items in inventory, but got %v", itemCount)
	}
	fmt.Println("Test for inventoryRandomizer function passed")
}
