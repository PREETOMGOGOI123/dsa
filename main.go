package main

import (
	"dsa/plusOne"
	"dsa/sqrt"
	"fmt"
)

func main() {
	input := []int{1, 2, 3}
	result := plusOne.PlusOne(input)

	fmt.Println(result)

	sqrt := sqrt.MySqrt(130)
	fmt.Println(sqrt)
}
