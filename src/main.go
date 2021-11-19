package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//Simulacion de contructor
func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: true,
	}
}

func (this *Employee) setId(id int) {
	this.id = id

}

func (this *Employee) setName(name string) {
	this.name = name

}

func (this *Employee) getId() int {
	return this.id

}
func (this *Employee) getName() string {
	return this.name

}

func main() {
	employee := Employee{}
	fmt.Printf("%v", employee)

	employee.id = 4
	employee.name = "Facundo"
	fmt.Println(employee)

	employee.setId(11)
	employee.setName("Facundo")
	fmt.Println(employee)

	fmt.Println(employee.getId())

	//Puntero
	employee1 := new(Employee)
	fmt.Println(*employee1)

	//Contructor
	employee2 := NewEmployee(1, "Facundo")
	fmt.Println(*employee2)
}
