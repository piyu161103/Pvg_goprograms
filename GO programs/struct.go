package main

import "fmt"
type person struct{
	name string
	age  int
}

func main() {
p := person{name: "sakshi",age: 20}
    fmt.Println("person details:",p)
}
