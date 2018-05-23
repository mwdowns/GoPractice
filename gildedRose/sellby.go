package main

func (i *item) anotherDayCloserToDeath() {
	i.SellBy = i.SellBy - 1
}

func (i *item) cursed() {
	i.SellBy = i.SellBy - 1
}
