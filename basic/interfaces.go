package main

import (
	"fmt"
)

var pl = fmt.Println

//  in the words of Derek - "Interface are like contracts - that say, if anything inherits it, that they must implement some defined methods"

type Animal interface {
	AnrgySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AnrgySound() {
	pl("hissss")
}

func (c Cat) HappySound() {
	pl("meoooowww")
}

func main() {
	var kitty Animal
	kitty = Cat("kitty")

	kitty.AnrgySound()
	kitty.HappySound()
	var kitty2 Cat = kitty.(Cat)

	kitty2.Attack()
	pl("the names of kitty is: ", kitty2.Name())
}
