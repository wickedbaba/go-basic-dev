package main

import (
	"fmt"
)

var pl = fmt.Println

// ----------------------------------------------------------------------------------------------------------------------------
// defined types
type ML float64
type DL float32

// ----------------------------------------------------------------------------------------------------------------------------

//  creating generics

type MyConstraint interface {
	int | float64
}

// ----------------------------------------------------------------------------------------------------------------------------

//  structs

type player struct {
	name string
	age  int
	bal  float64
}

func getSumGen[T MyConstraint](x T, y T) (z T) {

	return x + y
}

type rectangle struct {
	length, height float64
}

// functions for struct
func (r rectangle) getArea() float64 {
	return r.length * r.height
}

// ----------------------------------------------------------------------------------------------------------------------------
// composition
// golang follows composition and not inheritance

type contact struct {
	fname string
	lname string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("The info of the business - %s -  is : %s %s ", b.name, b.contact.fname, b.contact.phone)
}

func main() {

	pl(getSumGen(5, 4))
	pl(getSumGen(5.6, 4.9))

	abc := rectangle{23.0, 27.86}

	pl(abc.getArea())

	contactOne := contact{
		"James",
		"Jordan",
		"9876543210",
	}

	businessOne := business{
		"Tech Industries",
		"sklark industries",
		contactOne,
	}

	businessOne.info()

}
