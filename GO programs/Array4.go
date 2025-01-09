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

    // Bubble Sort
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    fmt.Println("Array in ascending order:")
    for i := 0; i < n; i++ {
        fmt.Println(arr[i])
    }
}