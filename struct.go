package main

import (
	"fmt"
)

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {
// 	var employee1 Employee
// 	employee1.name = "John"
// 	employee1.age = 30
// 	employee1.division = "IT"

// 	var employee2 = Employee{name: "Ananda", age: 20, division: "Finance"}

// 	// fmt.Println("Employee name:", employee.name)
// 	// fmt.Println("Employee age:", employee.age)
// 	// fmt.Println("Employee division:", employee.division)

// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

//pointer to a struct

// func main() {
// 	var employee1 = Employee{name: "John", age: 30, division: "IT"}
// 	var employee2 *Employee = &employee1

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)

// 	fmt.Println(strings.Repeat("#", 20))

// 	employee2.name = "Ananda"

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)
// }

//embedded struct

// type Person struct {
// 	name string
// 	age  int
// }

// type Employee struct {
// 	division string
// 	person   Person
// }

// func main() {
// 	var employee1 = Employee{}

// 	employee1.person.name = "John"
// 	employee1.person.age = 30
// 	employee1.division = "IT"

// 	fmt.Printf("Employee1: %+v\n", employee1)
// }

//anonymous struct

// func main() {
// 	//anonymous struct without field name
// 	var employee1 = struct {
// 		person   Person
// 		division string
// 	}{}
// 	employee1.person = Person{name: "John", age: 30}
// 	employee1.division = "IT"

// 	//anonymous struct with field name
// 	var employee2 = struct {
// 		person   Person
// 		division string
// 	}{
// 		person:   Person{name: "Ananda", age: 20},
// 		division: "Finance",
// 	}
// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

//slice with struct

type Person struct {
	name string
	age  int
}

func main() {
	var people = []Person{
		{name: "John", age: 30},
		{name: "Ananda", age: 20},
		{name: "Jane", age: 25},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
