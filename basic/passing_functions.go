package main

import (
	"fmt"
)

var pl = fmt.Println

func userFunc(f func(int, int) int, x int, y int) int {
	return f(x, y)
}

func sumVal(x int, y int) int {
	return x + y
}

func main() {
	intSum := func(x int, y int) int { return x + y }

	pl("The sum is : ", intSum(5, 4))

	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}

	changeVar()
	pl("New Value:", samp1)

	pl("Output: ", userFunc(sumVal, 5, 4))

}
