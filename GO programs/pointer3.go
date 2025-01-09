package main

import "fmt"

func increment(num *int) {
    *num++
}

func main() {
    var num int = 10
    fmt.Println("Before increment:", num)

    increment(&num)
    fmt.Println("After increment:", num)
}