package main

import (
    "fmt"
    "strconv"
)
var something string = "hello"
const something2 = "world"

func main()  {
    
    // data := cityInfo{"hello", "cleveland"}
    // message1 := printMessage(data.greeting, data.name)
    // fmt.Println(message1)
    
    // winPer,city := recordInfo(2, 16, data)
    // fmt.Printf("the %v football team had a %v winning percentage last year.\n", city, winPer)
    
    city1 := cityInfo{"hello", "cleveland", 2, 16}
    city2 := cityInfo{"good morning", "atlanta", 15, 19}
    city3 := cityInfo{"fuck off", "boston", 16, 19}
    
    helloCity(city1, city2, city3)
    fmt.Println("\n")
    cities :=[]cityInfo { city1, city2, city3 }
    helloCity(cities...)
    
}

type cityInfo struct {
    greeting string
    name string
    wins float32
    gamesPlayed float32
}

func printMessage(var1 string, var2 string) string {
    return var1 + " " + var2 + "!"
}

func recordInfo(city cityInfo) (message2 string) {
    winPer := city.wins / city.gamesPlayed
    var per = strconv.FormatFloat(float64(winPer), 'f', 3, 32)
    return "the " + city.name + " football team had a " + per + " winning percentage last year."
}

func helloCity(cities ...cityInfo) {
    for _,city := range cities {
        message := printMessage(city.greeting, city.name)
        fmt.Println(message)
        message2 := recordInfo(city)
        fmt.Println(message2)
    }
}