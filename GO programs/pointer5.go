package main

import "fmt"

func swap(x, y *int) {
    temp := *x
    *x = *y
    *y = temp
}

func main() {
    var a, b int = 5, 10
    fmt.Println("Before swap: a =", a, "b =", b)

    swap(&a, &b)
    fmt.Println("After swap: a =", a, "b =", b)
}