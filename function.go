package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	// greet("Hactiv8", 6)
// 	// greet("Hactiv8", "Jl Thamrin")

// 	//function with return value
// 	var names = []string{"Hactiv8", "Jakarta"}
// 	var printMsg = greet("Hei", names)

// 	fmt.Println(printMsg)
// }

// func greet(name string, age int8) {
// 	fmt.Printf("Hello there! My name is %s and I am %d years old.", name, age)
// }

// func greet(name, address string) {
// 	fmt.Println("Hello there! My name is", name, "and I am from", address)
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }

// func main() {
// 	var diameter float64 = 15

// 	var area, circumfreence = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference:", circumfreence)
// }

// func calculate(d float64) (float64, float64) {
// 	//count area
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	//count circumference
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	//count area
// 	area = math.Pi * math.Pow(d/2, 2)
// 	//count circumference
// 	circumference = math.Pi * d

// 	return
// }

// func main() {
// 	studentLists := print("Hactiv8", "Kidz", "Code")
// 	fmt.Printf("%v", studentLists)
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student %d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

// func main() {
// 	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	result := sum(numberList...)

// 	fmt.Println("Result:", result)
// }

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, v := range numbers {
// 		total += v
// 	}
// 	return total
// }

func main() {
	profile("Sushi", "Steak", "Pizza", "Pasta")
}

func profile(name string, faveFoods ...string) {
	mergeFavefoods := strings.Join(faveFoods, ", ")

	fmt.Println("Hello there, im", name)
	fmt.Println("I really love to eat ", mergeFavefoods)
}
