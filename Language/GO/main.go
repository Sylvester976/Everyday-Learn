package main

import (
	"fmt"
	"strings"
)

func main() {
	//Standard Libraries
	salutations := "Hello World"
	fmt.Println(salutations, strings.Contains(salutations, "Hello"))
	fmt.Println(salutations, strings.ReplaceAll(salutations, "Hello", "Salut"))
	fmt.Println(salutations, strings.ToUpper(salutations))
}
