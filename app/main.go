package main

import (
	stuff "example/project/mypackages"
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {

	pl("Checking package reach ... ", stuff.Name)

	intArr := []int{1, 2, 3, 4, 5}

	strArr := stuff.IntArrToStr(intArr)

	pl("the result: ", strArr)
	pl("the type of result: ", reflect.TypeOf(strArr))
}
