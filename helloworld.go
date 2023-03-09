package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// a variable which refers to println
var pl = fmt.Println

func main() {

	pl("Hello World from Go")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("hello", name)
	} else {
		log.Fatal(err)
	}

}
