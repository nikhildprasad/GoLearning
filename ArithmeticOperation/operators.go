package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

func main() {
	fmt.Println("Arithmetic Operations:")
	fmt.Println("Addition:", Add(5, 3))
	fmt.Println("Subtraction:", Subtract(5, 3))
	fmt.Println("Multiplication:", Multiply(5, 3))
	fmt.Println("Division:", Divide(5, 3))
}
