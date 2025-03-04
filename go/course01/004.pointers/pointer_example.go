package main

import "fmt"

//v1
// func main() {
// 	age := 32
// 	var agePointer *int
// 	agePointer = &age

// 	fmt.Println("Age: ", *agePointer)

// 	adultYears := getAdultYears(agePointer)
// 	fmt.Println("adultYears: ", adultYears)
// }

// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// v2
func main() {
	age := 32

	fmt.Println("Age: ", age)

	getAdultYears(&age)

	fmt.Println("adultYears: ", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
