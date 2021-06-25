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

	// Arrays
	grades := [...]int { 97, 85, 73, 100 }
	fmt.Println(grades)

	var students [3]string
	fmt.Println(students)
	students[0] = "Kevin"
	students[2] = "Gautier"
	students[1] = "Kayisire"
	fmt.Println(len(students))
	fmt.Println(students)
}