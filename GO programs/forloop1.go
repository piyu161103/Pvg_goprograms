package main

import "fmt"

func main() {
    var num, temp, remainder, result int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    temp = num
    for temp != 0 {
        remainder = temp % 10
        result += remainder * remainder * remainder
        temp /= 10
    }

    if result == num {
        fmt.Println(num, "is an Armstrong number.")
    } else {
        fmt.Println(num, "is not an Armstrong number.")
    }
}