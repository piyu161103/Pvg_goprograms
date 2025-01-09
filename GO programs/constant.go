package main

import "fmt"
const pi float64=3.5678
const language string="Go"

func main() {
	var radius float64=6.0
	var area float64=pi* radius*radius
    fmt.Println("constant pi:",pi)
	fmt.Println("constantn language:",language)
	fmt.Println("radius:",radius)
	fmt.Println("area of the circle:",area)
}
