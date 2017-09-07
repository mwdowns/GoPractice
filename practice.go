package main

import "fmt"

var something string = "hello"
const something2 = "world"

func main()  {
    
    data := message{"hello", "cleveland"}
    message1 := printMessage(data.greeting, data.name)
    fmt.Println(message1)
    
    winPer,city := newFunction(2, 16, data)
    fmt.Printf("the %v football team had a %v winning percentage last year.\n", city, winPer)
    
}

type message struct {
    greeting string
    name string
}

func printMessage(var1 string, var2 string) string {
    return var1 + " " + var2 + "!"
}

func newFunction(wins float32, games float32, city message) (winPer float32, cityName string) {
    cityName = city.name
    winPer = wins / games
    return
}