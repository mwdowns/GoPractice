package main

import "fmt"

var something string = "hello"
const something2 = "world"

func main()  {
    
    data := message{"hello", "cleveland"}
    message1 := printMessage(data.greeting, data.name)
    fmt.Println(message1)
    
}

type message struct {
    greeting string
    name string
}

func printMessage(var1 string, var2 string) string {
    return var1 + " " + var2 + "!"
}