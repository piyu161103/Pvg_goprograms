package main

import (
	"fmt"
	"log"
)

func main() {
	var num int
	fmt.Println("enter an integer:")
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("you entered:", num)
}
