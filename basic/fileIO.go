package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	// // os.Create() function ->
	// // It creates a new file if it does not exist.
	// // It truncates the file if it exists.
	// file, err := os.Create("data.txt")

	// if err != nil {
	// 	// log.Fatal()
	// 	// prints specified message with timestamp on the console screen.
	// 	// The log.Fatal() is similar to the log.Print() function, but is followed by a call to os.Exit function
	// 	log.Fatal(err)
	// }

	// //  used to close the file
	// // One of the primary uses of a defer statement is for cleaning up resources, such as open files, network connections, and database handles.
	// //  defer works by executing the statement after the function returns

	// // better definition ->
	// // A defer statement adds the function call following the defer keyword onto a stack.
	// // All of the calls on that stack are called when the function in which they were added returns.
	// // Because the calls are placed on a stack, they are called in last-in-first-out order.

	// defer file.Close()

	// iPrimeArr := []int{2, 3, 5, 7, 11}

	// var sPrimeArr []string

	// for _, num := range iPrimeArr {
	// 	sPrimeArr = append(sPrimeArr, strconv.Itoa(num))
	// }

	// //  this is one of the reasons why I can call file.WriteString here
	// for _, num := range sPrimeArr {
	// 	// WriteString -> writes the contents of a string to a file
	// 	_, err := file.WriteString(num + "\n")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// file, err = os.Open("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scan1 := bufio.NewScanner(file)

	// for scan1.Scan() {
	// 	pl("Prime : ", scan1.Text())
	// }

	// // for the whole scan operation above, it goes through all the errors that occured; and if it is not equal to nil, then prints it.
	// if err := scan1.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
