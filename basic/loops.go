package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	//  simple for loop with decrements
	for x := 5; x > 0; x-- {
		pl(x)
	}

	//  for loop being used as a while loop
	index := 5

	for index > 0 {
		pl(index, "here")
		index--

	}
}
