package main

import "fmt"

func main(){
	val:=sum(5,10)
	fmt.Println("sum is ",val)
}
func sum(a int, b int )int{
	fmt.Println("hello")
	return a+b
}