package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var Pl = fmt.Println

func main() {

	// var nameOfVariable type_of_variable

	// NOTE: if name of variable start with capital, then it is a global variable which can be accessed from outside

	//  main data types
	// 1. int - 0
	// 2. float64 - 0.0
	// 3. bool - false
	// 4. string - ""
	// 5. rune - ""
	Pl("Whats my name")

	//  to see datatype ->
	Pl(reflect.TypeOf(25))

	//  type casting ->

	var1 := 1.5
	var2 := int(var1)

	Pl(reflect.TypeOf(var2))
	// string to int
	var3 := "509090"
	var4, err := strconv.Atoi(var3) // means ascii to integer

	if err == nil {
		Pl(var4)
	} else {
		Pl(err)
	}

	// int to string

	var5 := 590909
	var6 := strconv.Itoa(var5) // means integer to ascii

	Pl(var5, var6, reflect.TypeOf(var6))

	// string to float
	var7 := "100"
	var8, err := strconv.ParseFloat(var7, 64)

	Pl(var7, var8, err, reflect.TypeOf(var8))

	//  float to string

	var9 := 100.10
	var10 := fmt.Sprintf("%f", var9)

	Pl(var10)
}
