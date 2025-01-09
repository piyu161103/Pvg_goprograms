package main

import "fmt"

func main() {
    var num, i int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    i = 1
    for {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
        i++
        if i > 10 {
            break
        }
    }
}
