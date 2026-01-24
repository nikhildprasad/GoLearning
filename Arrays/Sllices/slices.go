package main

import "fmt"

func main() {
	var a [5]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "from"
	a[3] = "Go"
	a[4] = "Lang"

	slice := a[1:4]

	fmt.Println("Array:", a)
	fmt.Println("Slice:", slice)
}
