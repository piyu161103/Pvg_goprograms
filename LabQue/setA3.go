package main 
import "fmt"
func main(){
	var a,b int
	fmt.Print("Enter two numbers(a and b):")
	fmt.Scan(&a,&b)
	fmt.Println("before swap: a=",a,",b=",b)
    a = a + b
    b = a - b
    a = a - b

    fmt.Println("After Swap: a =", a, ", b =", b)
}