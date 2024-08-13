package main

import "fmt"

func main() {
	//Arrays and Slices
	var ages [3]int = [3]int{40, 20, 30}
	var new_age = [3]int{40, 20, 30}
	fmt.Println(ages, new_age, len(ages), len(new_age))

	names := [3]string{"My", "Name", "Is"}
	fmt.Println(names)

	//slices in Go
	scores := []int{20, 21, 22}
	fmt.Println(scores, len(scores))
	scores[2] = 87
	fmt.Println(scores)
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

}
