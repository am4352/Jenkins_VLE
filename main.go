
package main

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Go CI Pipeline - Jenkins")

	result := Multiply(5, 4)
	fmt.Println("Multiplication result:", result)
}
