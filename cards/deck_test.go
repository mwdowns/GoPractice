package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	fmt.Println("len test passed")
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts as first card, got %v", d[0])
	}
	fmt.Println("First card test passed")
	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected Ten of Clubs as last card, got %v", d[len(d)-1])
	}
	fmt.Println("Last card test passed")
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d2 := newDeck()
	d2.shuffle()
	count := 0
	for i := range d {
		if d[i] == d2[i] {
			count++
		}
	}
	if count > 25 {
		t.Errorf("Not very random. %v matches", count)
	}
	fmt.Println("Ranomizer test passed")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	num := 5
	hand, remainingDeck := deal(d, num)
	if len(hand) != num {
		t.Errorf("Expected hand of 5 cards, got %v", len(hand))
	}
	fmt.Println("Hand len test passed")
	if len(remainingDeck) != len(d)-num {
		t.Errorf("Expected remainingDeck to be 47, got %v", len(remainingDeck))
	}
	fmt.Println("Remaining deck len test passed")
}
