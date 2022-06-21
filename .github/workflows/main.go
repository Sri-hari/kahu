package main

import "fmt"
func main() {
    a := "hello world"
    b := false
    
    fmt.Println("a - " a)
    fmt.Println ("address - ", &a)
    fmt.Printf("%%v - %v\n", a)
    fmt.Printf("%%T - %T\n", b)
}
 
