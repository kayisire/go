package main

import (
	"fmt"
	"strconv"
)

var (
	someString string
	someInteger int
)

func main() {
	fmt.Println("Hello World!")
	
	someString = "what it is!"
	someInteger = 30

	fmt.Println(someString)
	fmt.Println(someInteger)

	someString = "shadowing, woo!"
	fmt.Println(someString)

	someImplicitString := "what it do!"
	someImplicitInteger := 10

	fmt.Println(someImplicitString)
	fmt.Println(someImplicitInteger)

	i := 10
	var j string = strconv.Itoa(i)

	fmt.Println(j)
}