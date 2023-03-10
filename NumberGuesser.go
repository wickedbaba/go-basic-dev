package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {

	// use time to generate a random seed -> then generate a random number
	//  take input from the user and check if the user can guess the number
	// for loop -> input -> convert to int -> compare -> use if else to give hints

	randomSeed := time.Now().Unix()
	rand.Seed(randomSeed)
	randomNumber := rand.Intn(50) + 1

	for true {
		pl("Guess a number between 1 to 50 inclusive:")

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')

		if err != nil {
			pl("fatal error", err)
		}

		guess = strings.TrimSpace(guess)
		intGuess, err := strconv.Atoi(guess)

		if err == nil {

			if intGuess == randomNumber {
				pl("You are correct, the random number is :", guess)
				break
			} else if intGuess > randomNumber {
				pl("you're a bit higher")
				continue
			} else {
				pl("you're a bit lower")
				continue
			}
		} else {
			pl("We have an error:", err)
		}

	}

}
