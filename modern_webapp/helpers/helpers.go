package helpers

import (
	"math/rand"
	"time"
)

type Animal interface {
	SayName() string
	Says() string
	NumberOfLegs() int
}

type Pet struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Word string `json:"word"`
	Legs int    `json:"legs"`
}

func (p *Pet) SayName() string {
	return p.Name
}

func (p *Pet) Says() string {
	return p.Word
}

func (p *Pet) NumberOfLegs() int {
	return p.Legs
}

func GenerateRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
