package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n) // Create a slice (1D array)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Println("The array elements are:")
    for i := 0; i < n; i++ {
        fmt.Println(arr[i])
    }
}