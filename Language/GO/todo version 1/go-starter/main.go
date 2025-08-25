package main

import (
	"fmt"
	"net/http"
	"strings"
)

// global variable
var taskOne = "Complete GoLang project"
var taskTwo = "Read GoLang documentation"
var taskThree = "Practice coding challenges"
var taskItems = []string{taskOne, taskTwo, taskThree}

func Introduction(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, " ##### Welcome to my Todolist App! #####")
}

func PrintTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, strings.ToUpper("List of my Todos"))
	for index, item := range taskItems {
		fmt.Fprintf(writer, "%d: %s\n", index+1, item)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}

func main() {
	http.HandleFunc("/", Introduction)
	http.HandleFunc("/showtasks", PrintTasks)
	http.ListenAndServe(":8080", nil)
}

//These were used to make me learn golang

//package main
//
//import (
//"fmt"
//"strings"
//)
//
//func TodoStarterPack() {
//	fmt.Println(" ##### Welcome to my Todolist App! #####")
//	fmt.Println()
//}
//
//func title1() {
//	fmt.Println("Introduction ")
//	fmt.Println()
//
//}
//
//func title2() {
//	fmt.Println(strings.ToUpper("List of my Todos"))
//}
//
//func title3() {
//	fmt.Println()
//	fmt.Println("My Variables")
//}
//
//func title4() {
//	fmt.Println(strings.ToUpper("My Variables 2"))
//}
//
//func nameDisplay() {
//	fmt.Println("1. My name is Sylvester Otieno")
//	fmt.Println()
//}
//
//func myVariables() {
//	var no1 = 1
//	var no2 = 2
//	var no3 = 3
//	var taskOne = " \n 1. Complete GoLang project \n"
//	var taskTwo = " \n 2. Read GoLang documentation \n "
//	var taskThree = "\n 3. Practice coding challenges \n "
//	fmt.Println(taskOne, taskTwo, taskThree)
//	fmt.Println("Total number of tasks:", no1+no2+no3)
//	fmt.Println()
//}
//
//func myVariables2() {
//	taskOne := " \n 1. Complete GoLang project \n"
//	taskTwo := " \n 2. Read GoLang documentation \n "
//	taskThree := "\n 3. Practice coding challenges \n "
//	fmt.Println(taskOne, taskTwo, taskThree)
//	fmt.Println()
//	totalTasks := 3
//	fmt.Println("Total number of tasks:", totalTasks)
//	fmt.Println()
//}
//
//func myList() {
//	taskOne := " \n 1. Complete GoLang project \n"
//	taskTwo := " \n 2. Read GoLang documentation \n "
//	taskThree := "\n 3. Practice coding challenges \n "
//	numbers := []int{1, 2, 3}
//	tasks := []string{taskOne, taskTwo, taskThree}
//	fmt.Println("My Tasks:", tasks)
//	fmt.Println("Total number of tasks:", len(numbers))
//	fmt.Println()
//}
//
//func myInterface() {
//	taskOne := " \n 1. Complete GoLang project \n"
//	numberOne := 1
//	myItems := []interface{}{taskOne, numberOne}
//	fmt.Println("My Items:", myItems)
//	fmt.Println("Total number of items:", len(myItems))
//	fmt.Println()
//}
//func loopTitle() {
//	fmt.Println("My Loops")
//	fmt.Println()
//}
//
//func MyLoops() {
//	taskOne := " Complete GoLang project "
//	taskTwo := " Read GoLang documentation"
//	taskThree := " Practice coding challenges "
//	numberOne := 1
//	numberTwo := 2
//	numberThree := 3
//
//	myItems := []interface{}{taskOne, taskTwo, taskThree, numberOne, numberTwo, numberThree}
//	//for ranger loop
//	for index, item := range myItems {
//		fmt.Printf("%d: %v\n", index+1, item)
//	}
//
//	fmt.Println()
//	fmt.Println()
//
//	for index, item := range myItems {
//		fmt.Println(index, item)
//	}
//	fmt.Println()
//	fmt.Println()
//
//	//for range loop 2
//	for _, item := range myItems {
//		fmt.Println(item)
//	}
//	fmt.Println()
//	fmt.Println()
//	//for loop
//	for i := 0; i < len(myItems); i++ {
//		fmt.Println(myItems[i])
//	}
//	fmt.Println()
//	fmt.Println()
//}
//func valueFunction() {
//	fmt.Println("This is a value function")
//	fmt.Println()
//}
//
//func functionValue(val int) {
//	if val == 1 {
//		fmt.Println("Value is one")
//	} else {
//		fmt.Println("Value is not one")
//	}
//}
//
//func main() {
//	TodoStarterPack()
//	title1()
//	nameDisplay()
//	title2()
//	title3()
//	myVariables()
//	title4()
//	myVariables2()
//	myList()
//	myInterface()
//	loopTitle()
//	MyLoops()
//	valueFunction()
//	functionValue(1)
//}
