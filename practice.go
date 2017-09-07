package main

import "fmt"

var something string = "hello"
const something2 = "world"

func main()  {
    
    message1 := message{greeting: "hello", person: "cleveland"}
    
    fmt.Println(message1.greeting, message1.person + "!")
}

type message struct {
    greeting string
    person string
}