package main

import "fmt"

// func main() {
// 	var firstNumber *int
// 	var secondNumber *int

// 	_, _ = firstNumber, secondNumber
// }

// func main() {
// 	var firstNumber int = 4
// 	var secondNumber *int = &firstNumber

// 	fmt.Println("firstNumber (value):", firstNumber)
// 	fmt.Println("secondNumber (memory address):", &firstNumber)

// 	fmt.Println("secondNumber (value):", *secondNumber)
// 	fmt.Println("secondNumber (memory address):", secondNumber)
// }

// func main() {
// 	var firstPerson string = "John"
// 	var secondPerson *string = &firstPerson

// 	fmt.Println("firstPerson (value):", firstPerson)
// 	fmt.Println("secondPerson (memory address):", &firstPerson)
// 	fmt.Println("secondNumber (value):", *secondPerson)
// 	fmt.Println("secondNumber (memory address):", secondPerson)

// 	fmt.Println("\n", strings.Repeat("#", 40), "\n")

// 	*secondPerson = "Doe"

// 	fmt.Println("firstPerson (value):", firstPerson)
// 	fmt.Println("secondPerson (memory address):", &firstPerson)
// 	fmt.Println("secondNumber (value):", *secondPerson)
// 	fmt.Println("secondNumber (memory address):", secondPerson)
// }

func main() {
	var a int = 10
	fmt.Println("Before: ", a)
	changeValue(&a)
	fmt.Println("After: ", a)
}

func changeValue(number *int) {
	*number = 20
}
