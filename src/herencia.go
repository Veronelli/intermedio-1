package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.id = 1

	ftEmployee.name = "Facundo"
	ftEmployee.age = 18

	fmt.Println(ftEmployee)

}
