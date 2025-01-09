package main

import "fmt"

func main() {
    var str1, str2 string

    fmt.Print("Enter the first string: ")
    fmt.Scan(&str1)

    fmt.Print("Enter the second string: ")
    fmt.Scan(&str2)

    if str1 == str2 {
        fmt.Println("The strings are equal.")
    } else {
        fmt.Println("The strings are not equal.")
    }
}