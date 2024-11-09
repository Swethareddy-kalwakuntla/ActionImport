package main

import "fmt"

func main(){
	val:=sum(5,10)
	fmt.Println(val)
}
func sum(a int, b int )int{
	return a+b
}