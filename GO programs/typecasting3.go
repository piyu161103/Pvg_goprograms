package main

import (
	"fmt"
       "strconv"
)
func main() {
	var str string="234"
	var num int
	var err error
	num,err=strconv.Atoi(str)
	if err!=nil{
		fmt.Println("error:",str)
	} else{
		fmt .Println("string value:",str)
	
    fmt.Println("converted int values:",num)
}
}