package main

import "fmt"

func main() {
    var a, b, choice int
    fmt.Println("1. Add")
    fmt.Println("2. Subtract")
    fmt.Println("3. Multiply")
    fmt.Println("4. Divide")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    switch choice {
    case 1:
        fmt.Println("Result:", a+b)
    case 2:
        fmt.Println("Result:", a-b)
    case 3:
        fmt.Println("Result:", a*b)
    case 4:
        if b != 0 {
            fmt.Println("Result:", a/b)
        } else {
            fmt.Println("Division by zero is not allowed.")
        }
    default:
        fmt.Println("Invalid choice.")
    }
}