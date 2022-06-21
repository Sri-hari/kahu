package main

import (
    "flag"
    "fmt"
)

var sayHi bool

func init() {
    flag.BoolVar(&sayHi, "sayHi", false, "Say hi or not")
    flag.Parse()
}

func main() {
    if sayHi {
        fmt.Println("Hi!")
    }
}
