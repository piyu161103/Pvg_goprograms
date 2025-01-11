package main

import "fmt"

func main() {
    var str1, str2 string

    fmt.Print("Enter the first string: ")
    fmt.Scan(&str1)

    fmt.Print("Enter the second string: ")
    fmt.Scan(&str2)
    ptr1 := &str1
    ptr2 := &str2
    concatenated := *ptr1 + *ptr2

    fmt.Println("Concatenated string:", concatenated)
}
