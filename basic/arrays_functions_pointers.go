package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func getSum(x int, y int) int {

	return x + y
}

func returnMultiple(x int, y int) (int, int) {
	return x * y, x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {

	if y == 0.0 {
		return 0, fmt.Errorf("You cant divide by zero")
	} else {
		return x / y, nil
	}
}

func unspecInput(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func passingArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func changeValue(myPtr *int) {
	*myPtr = 12

}

func main() {

	arrNums := []int{1, 2, 3}
	for index, num := range arrNums {
		pl(index, num)
	}

	//  initialise an empty array
	var emptyArr [5]int // -> of size 5
	emptyArr[0] = 51

	//  length
	pl(len(arrNums))

	//  2D array

	// var twoDArr [5][5]int
	// or
	twoDArr := [2][2]int{
		{1, 2},
		{3, 4},
	}
	pl(twoDArr)

	// ----------------------------------------------------------------------------------------------------------------------------

	// slicing

	// make -> allocates and initializes an object of type -> slice, map, chan (only)
	sl1 := make([]string, 6)

	sl1[0] = "Wuba"
	sl1[1] = "luba"
	sl1[2] = "dub"
	sl1[3] = "dub"
	sl1[4] = "dub"
	sl1[5] = "dub"

	pl(reflect.TypeOf(sl1))
	for index, sl := range sl1 {
		pl(index, sl)
	}

	// ----------------------------------------------------------------------------------------------------------------------------

	// functions

	pl(getSum(32, 57))

	// returning multiple values

	pl(returnMultiple(32, 57))

	// returning an error
	pl(getQuotient(5, 0))

	//  for an unspecified length of input
	pl(unspecInput(1, 2, 3, 4, 5, 5, 6, 6))

	// passing array
	varArr := []int{1, 2, 3, 4, 5, 5, 6, 6}
	pl(passingArray(varArr))
	// pl(passingArray([]int{1, 2, 3, 4, 5, 5, 6, 6}))

	// ----------------------------------------------------------------------------------------------------------------------------
	// pointers

	f3 := 10

	pl(f3)
	changeValue(&f3)
	pl(f3)

}
