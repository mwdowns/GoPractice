package main

import "fmt"

var something string = "hello"
const something2 = "world"

func main()  {
    
    // data := cityInfo{"hello", "cleveland"}
    // message1 := printMessage(data.greeting, data.name)
    // fmt.Println(message1)
    
    // winPer,city := newFunction(2, 16, data)
    // fmt.Printf("the %v football team had a %v winning percentage last year.\n", city, winPer)
    
    city1 := cityInfo{"hello", "cleveland", 2, 16}
    city2 := cityInfo{"good morning", "atlanta", 15, 19}
    city3 := cityInfo{"fuck off", "boston", 16, 19}
    
    helloCity(city1.name, city2.name, city3.name)
    fmt.Println("\n")
    cities :=[]string { city1.name, city2.name, city3.name }
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

func newFunction(city cityInfo) (winPer float32, cityName string) {
    cityName = city.name
    winPer = city.wins / city.gamesPlayed
    return
}

func helloCity(cities ...string) {
    for _,city := range cities {
        message := printMessage("hello", city)
        fmt.Println(message)
    }
}