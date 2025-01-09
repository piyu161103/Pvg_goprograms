package main

import "fmt"

func main() {
    var rows int
    fmt.Print("Enter the number of rows for Pascal's Triangle: ")
    fmt.Scan(&rows)

    triangle := make([][]int, rows)
    for i := 0; i < rows; i++ {
        triangle[i] = make([]int, i+1)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                triangle[i][j] = 1
            } else {
                triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
            }
            fmt.Printf("%d ", triangle[i][j])
        }
        fmt.Println()
    }
}