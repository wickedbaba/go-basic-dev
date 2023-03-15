package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var pl = fmt.Println

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

// error checker to handle errors
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// a writing to the browser using the writer
// the writer is used to - write to the browser,
// it creates a message and adds it to the response which displays it in the browser
func write(writer http.ResponseWriter, msg string) {

	_, err := writer.Write([]byte(msg)) // doing a type conversion to bytes

	errorCheck(err)
}

// request coming from the browser
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Internet")
}

func spanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola Internet")
}

func getStrings(fileName string) []string {
	var lines []string

	// open the file
	file, err := os.Open(fileName)

	//  checking for errors
	if os.IsNotExist(err) {
		return nil
	}

	errorCheck(err)

	defer file.Close()

	// scanning the file for text
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())

	// returning the string in the file
	return lines
}

// to get text information from todos.txt and display on screen
func interactHandler(writer http.ResponseWriter, request *http.Request) {

	// using function to get the string
	todoVals := getStrings("todos.txt")

	// testing purposes :
	// pl(todoVals)

	// creating a template using the html
	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)

	// create a todo list with the number
	todos := ToDoList{
		ToDoCount: len(todoVals),
		ToDos:     todoVals,
	}

	// todo: hoave to check how this works
	// Write the template to the ResponseWriter
	// Pass the todo struct data
	err = tmpl.Execute(writer, todos)

}

// used to handle new todo insertions
func newHandler(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("new.html")
	errorCheck(err)

	err = tmpl.Execute(writer, nil)
	errorCheck(err)
}

// used to store the new todo in the text file - "todos.txt"
func createHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)

	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)

	err = file.Close()
	errorCheck(err)

	http.Redirect(writer, request, "/interact", http.StatusFound)

}

// essentially, how the flow works is
// newHandler -> handles the new todo creation
//  createHandler -> stores the new todo in the "todos.txt" file
//  interactHandler -> goes through the "todos.txt" and displays the todos in it

func main() {

	// after running the app by doing - go run webapp.go
	// go to http://localhost:8080/hello  or http://localhost:8080/hola

	//  to interact with the todo, go to http://localhost:8080/interact

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)

	// this listens for browser requests and responds and
	// only receives a value if an error is thrown
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
