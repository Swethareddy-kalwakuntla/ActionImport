package main

import (
	"testing"
	"fmt"
)

func Test_Sum(t *testing.T){
	val:=sum(10,5)
	if val==15{
		fmt.Println("success")
	}else{
			t.Errorf("Sum(10, 5) = %d; want %d", val, 15)
	}
}