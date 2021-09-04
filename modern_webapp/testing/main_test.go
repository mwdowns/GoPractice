package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid", 100.00, 10.00, 10.00, false},
	{"invalid", 100.00, 0.00, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divideThings(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected error but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect error, but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// func TestHappy(t *testing.T) {
// 	_, err := divideThings(100.00, 10.00)
// 	if err != nil {
// 		t.Error("got error when I shouldn't have")
// 	}
// }

// func TestSad(t *testing.T) {
// 	_, err := divideThings(100.00, 0.00)
// 	if err == nil {
// 		t.Error("got an error when I should't have")
// 	}
// }
