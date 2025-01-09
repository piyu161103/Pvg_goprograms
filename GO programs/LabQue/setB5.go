package main

import "fmt"

func main() {
    ptr := new(int) // Allocates memory for an integer and returns a pointer to it
    fmt.Println("Address of allocated memory:", ptr)
    fmt.Println("Initial value at allocated memory:", *ptr)

    *ptr = 50 // Assign a value to the allocated memory
    fmt.Println("Updated value at allocated memory:", *ptr)
}