package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	// conditional operators : > < >= <= == !=
	// logical operators : && || !

	var age int

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err == nil {
		age, err = strconv.Atoi(userInput)
	} else {
		age = 0
		pl(err)
	}

	if (age >= 1) && (age <= 20) {
		pl("Young G")
	} else if age >= 21 && age <= 40 {
		pl("Adult G")
	} else {
		pl("Old G")
	}

}
