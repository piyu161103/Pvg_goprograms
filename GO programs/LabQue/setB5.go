package main

import "fmt"

func main() {
    ptr := new(int) 
    fmt.Println("Address of allocated memory:", ptr)
    fmt.Println("Initial value at allocated memory:", *ptr)

    *ptr = 50 
    fmt.Println("Updated value at allocated memory:", *ptr)
}
