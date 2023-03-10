package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {

	now := time.Now()

	pl(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//  we use time as a seed value to create a random value

	seedSeconds := now.Unix() // -> return number of seconds passed from january 1 1970
	rand.Seed(seedSeconds)

	pl(rand.Intn(50) + 1) // -> get value from 0 to 49
}
