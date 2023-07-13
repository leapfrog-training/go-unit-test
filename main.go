package main

import "fmt"

// Example function to be tested
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func main() {
	result := Add(2, 3)
	sresult := Subtract(10, 5)
	fmt.Println(result)
	fmt.Println(sresult)
}
