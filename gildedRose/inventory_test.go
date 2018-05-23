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
	fmt.Println("Test for newItem function passed.")
}

func TestNewInventory(t *testing.T) {
	i := newInventory()
	if reflect.TypeOf(i) != reflect.TypeOf(inventory{}) {
		t.Errorf("Expected i to be of inventory type struct, got %v", reflect.TypeOf(i))
	}
	fmt.Println("Test for newInventory function passed.")
}

func TestInventoryRandomizer(t *testing.T) {
	i := inventoryRandomizer(15)
	itemCount := len(i.EpicItems) + len(i.SpecialItems) + len(i.ConjuredItems) + len(i.NormalItems)
	if itemCount != 15 {
		t.Errorf("Expected 15 items in inventory, but got %v", itemCount)
	}
	fmt.Println("Test for inventoryRandomizer function passed.")
}

func TestAnotherDayCloserToDeath(t *testing.T) {
	i := newItem("New Item", 10, 10)
	i.anotherDayCloserToDeath()
	if i.SellBy != 9 {
		t.Errorf("Expected i.SellBy to equal 9, but got %v", i.SellBy)
	}
	fmt.Println("Test for AnotherDayCloserToDeath function passed.")
}

func TestDecreaseQuality(t *testing.T) {
	normalI := newItem("Normal Item", 10, 10)
	conjuredI := newItem("Conjured Item", 10, 10)
	normalI.decreaseQuality(false)
	conjuredI.decreaseQuality(true)
	if conjuredI.Quality != 8 {
		t.Errorf("Expected Conjured Item to have quality of 8, but got %v", conjuredI.Quality)
	}
	if normalI.Quality != 9 {
		t.Errorf("Expected Normal Item to have quality of 9, but got %v", normalI.Quality)
	}
	fmt.Println("Tests for decreaseQuality function passed")
}

func TestIncreaseQuality(t *testing.T) {
	specialI := newItem("Special Item", 10, 10)
	specialIWithFiveDaysLeft := newItem("Special Item With Five Days Left", 5, 10)
	specialIWithTwoDaysLeft := newItem("Special Item With Two Days Left", 2, 10)
	specialIWithNoDaysLeft := newItem("Special Item With No Days Left", 0, 10)
	agedBrie := newItem("Aged Brie", 10, 10)
	agedBrieAtZero := newItem("Aged Brie", 0, 10)
	agedBriePastSellBy := newItem("Aged Brie", -1, 10)
	specialI.increaseQuality()
	specialIWithFiveDaysLeft.increaseQuality()
	specialIWithTwoDaysLeft.increaseQuality()
	specialIWithNoDaysLeft.increaseQuality()
	agedBrie.increaseQuality()
	agedBrieAtZero.increaseQuality()
	agedBriePastSellBy.increaseQuality()
	if specialI.Quality != 11 {
		t.Errorf("Expected Special Item to have quality of 11, but got %v", specialI.Quality)
	}
	if specialIWithFiveDaysLeft.Quality != 12 {
		t.Errorf("Expected Special Item With Five Days Left to have quality of 12, but got %v", specialIWithFiveDaysLeft.Quality)
	}
	if specialIWithTwoDaysLeft.Quality != 15 {
		t.Errorf("Expected Special Item With Two Days Left to have quality of 15, but got %v", specialIWithTwoDaysLeft.Quality)
	}
	if specialIWithNoDaysLeft.Quality != 0 {
		t.Errorf("Expected Special Item With No Days Left to have quality of 0, but got %v", specialIWithNoDaysLeft.Quality)
	}
	if agedBrie.Quality != 11 {
		t.Errorf("Expected Aged Brie to have quality of 11, but got %v", agedBrie.Quality)
	}
	if agedBrieAtZero.Quality != 11 {
		t.Errorf("Expected Aged Brie At Zero to have quality of 11, but got %v", agedBrieAtZero.Quality)
	}
	if agedBriePastSellBy.Quality != 11 {
		t.Errorf("Expected Aged Brie Past SellBy to have quality of 11, but got %v", agedBriePastSellBy.Quality)
	}
	fmt.Println("Tests for increaseQuality function passed")
}
