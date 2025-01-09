package main

import "fmt"

func main() {
    var number int = 42
    var ptr1 *int = &number  // Pointer to the variable
    var ptr2 **int = &ptr1   // Pointer to the pointer

    fmt.Println("Value of number:", number)
    fmt.Println("Address of number:", &number)
    fmt.Println("Value of ptr1 (address of number):", ptr1)
    fmt.Println("Value pointed by ptr1:", *ptr1)
    fmt.Println("Value of ptr2 (address of ptr1):", ptr2)
    fmt.Println("Value pointed by ptr2 (value of ptr1):", *ptr2)
    fmt.Println("Value pointed by the pointer to pointer:", **ptr2)
}