package main

import "fmt"

func main() {
    var n, key int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Print("Enter the element to search for: ")
    fmt.Scan(&key)

    found := false
    for i := 0; i < n; i++ {
        if arr[i] == key {
            fmt.Println("Element found at index:", i)
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Element not found in the array.")
    }
}