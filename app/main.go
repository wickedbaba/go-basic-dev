package main

import (
	stuff "example/project/mypackages"
	"fmt"
	"log"
	"reflect"
)

var pl = fmt.Println

func main() {

	pl("Checking package reach ... ", stuff.Name)

	intArr := []int{1, 2, 3, 4, 5}

	strArr := stuff.IntArrToStr(intArr)

	pl("the result: ", strArr)
	pl("the type of result: ", reflect.TypeOf(strArr))

	// ----------------------------------------------------------------------------------------------------------------------------
	//  implementation of protected variables

	dateVar := stuff.Date{}

	err := dateVar.SetDay(01)

	if err != nil {
		log.Fatal(err)
	}

	err = dateVar.SetMonth(10)

	if err != nil {
		log.Fatal(err)
	}

	err = dateVar.SetYear(2007)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Todays date is : %d-%d-%d", dateVar.GetDay(), dateVar.GetMonth(), dateVar.GetYear())

}
