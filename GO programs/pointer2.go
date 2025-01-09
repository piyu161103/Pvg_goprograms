package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    var ptr *int

    for i := 0; i < len(arr); i++ {
        ptr = &arr[i]
        fmt.Printf("Element %d: %d\n", i, *ptr)
    }
}