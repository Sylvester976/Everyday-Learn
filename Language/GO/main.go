package main

import "fmt"

func main() {
	fmt.Println("Hello World, it's time to learn go")

	//Variables
	// Strings

	var member_name string = "Mr Golf Member"
	var executiveMemberName = "Mr CEO"
	var SeniorBoardMember string
	fmt.Println(member_name, executiveMemberName, SeniorBoardMember)

	//Overwrite the variable
	member_name = "Mr New Member"
	executiveMemberName = "Mr CTO"
	SeniorBoardMember = "Mr President"
	fmt.Println(member_name, executiveMemberName, SeniorBoardMember)

	//Declare Variables without keyword 'Var'(this cant be used outside the function, Used for New Variables)
	member_name2 := "Mr Otieno"
	executiveMemberName2 := "Mr Nyabenge"
	SeniorBoardMember2 := "Mr Sylvester"
	fmt.Println(member_name2, executiveMemberName2, SeniorBoardMember2)

	//End of String Variables
}
