package main

import "fmt"

func main() {
    var choice int
    fmt.Println("Menu:")
    fmt.Println("1. Pizza")
    fmt.Println("2. Burger")
    fmt.Println("3. Pasta")
    fmt.Println("4. Exit")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("You chose Pizza.")
    case 2:
        fmt.Println("You chose Burger.")
    case 3:
        fmt.Println("You chose Pasta.")
    case 4:
        fmt.Println("Exiting...")
    default:
        fmt.Println("Invalid choice.")
    }
}