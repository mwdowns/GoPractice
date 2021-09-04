package helpers

type Animal interface {
	SayName() string
	Says() string
	NumberOfLegs() int
}

type Pet struct {
	Type string `json:"type"`
	Name string
	Word string
	Legs int
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
