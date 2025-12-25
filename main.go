package main

import (
	"fmt"

	"anas.com/learn-go/helloworld"
)
import "anas.com/learn-go/integers"

func main() {
	greet := helloworld.Hello("Anas", "Spanish")
	println(greet)

	sum := integers.Add(5000, 2123123)
	fmt.Printf("sum = %d \n", sum)

	arr1 := make([]int, 10, 15)
	for i, v := range arr1 {
		fmt.Printf("arr1[%d] = %d \n", i, v)
	}
}
