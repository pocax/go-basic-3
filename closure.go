package main

import "fmt"

// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int
// 		for _, v := range numbers {
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}

// 	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println(evenNumbers(numbers...))
// }

//palindrome > IIFE (Immediately Invoked Function Expression)
// func main() {
// 	var isPalindrome = func(str string) bool {
// 		var temp string = ""
// 		for i := len(str) - 1; i >= 0; i-- {
// 			temp += string(byte(str[i]))
// 		}
// 		return temp == str
// 	}("katak")
// 	fmt.Println(isPalindrome)
// }

//closure as return value
// func main() {
// 	var studentLists = []string{"Marc", "Roz", "Jezz"}

// 	var find = findStudent(studentLists)

// 	fmt.Println(find("Marc"))
// }

// func findStudent(students []string) func(string) string {
// 	return func(s string) string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i
// 				break
// 			}
// 		}
// 		if student == "" {
// 			return fmt.Sprintf("Student %s not found", s)
// 		}
// 		return fmt.Sprintf("Student %s found at position %d", s, position+1)
// 	}
// }

//closure callbacks
func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println("Total Odd Numbers:", find)
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
