package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    max, min := arr[0], arr[0]
    for i := 1; i < n; i++ {
        if arr[i] > max {
            max = arr[i]
        }
        if arr[i] < min {
            min = arr[i]
        }
    }

    fmt.Println("Maximum element:", max)
    fmt.Println("Minimum element:", min)
}