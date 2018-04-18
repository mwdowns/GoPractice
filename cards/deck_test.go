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
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d2 := newDeck()
	d2.shuffle()
	count := 0
	for i := range d {
		if d[i] == d2[i] {
			count += 1
		}
	}
	fmt.Println(count)
	if count > 25 {
		t.Errorf("Not very random. %v matches", count)
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	num := 5
	hand, remainingDeck := deal(d, num)
	if len(hand) != num {
		t.Errorf("Expected hand of 5 cards, got %v", len(hand))
	}
	if len(remainingDeck) != len(d)-num {
		t.Errorf("Expected remainingDeck to be 47, got %v", len(remainingDeck))
	}
}
