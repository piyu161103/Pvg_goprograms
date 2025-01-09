package main

import "fmt"

func main() {
    var num, sum int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    for {
        sum += num % 10
        num /= 10
        if num == 0 {
            break
        }
    }
    fmt.Println("Sum of digits is", sum)
}