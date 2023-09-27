package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	name     string
	age      int
	position string
}

func main() {
	// Struct
	var employee1 Employee
	employee1.name = "Danang"
	employee1.age = 27
	employee1.position = "BE"

	fmt.Printf("Employee 1: %v\n", employee1)

	var employee2 *Employee = &employee1
	fmt.Printf("Employee 2: %v\n", employee2)

	employee2.position = "Fullstack"
	fmt.Printf("Employee 1: %v\n", employee1)
	fmt.Printf("Employee 2: %v\n", employee2)

	var employee3 = struct {
		name     string
		age      int
		position string
	}{
		name:     "Danang",
		position: "FE",
	}
	fmt.Printf("Employee 3: %v\n", employee3)

	// Function
	fmt.Println(sum(3, 4))
	mult, div := calculate(15, 3)
	fmt.Println(mult)
	fmt.Println(div)
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	nums := []int{5, 10, 15, 20}
	fmt.Println(sumAll(nums...))

	// Method
	student := Person{"Danang", 27}
	fmt.Println(student.Introduce("Hello!"))

	fmt.Println(student)
	student.ChangeName1()
	fmt.Println(student)
	student.ChangeName2()
	fmt.Println(student)

	// Reflect
	value := reflect.ValueOf(student)
	tipe := reflect.TypeOf(student)
	fmt.Println("Nilai:", value)
	fmt.Println("Tipe data:", tipe)
	value2 := reflect.ValueOf(student.name)
	tipe2 := reflect.TypeOf(student.name)
	fmt.Println("Nilai:", value2)
	fmt.Println("Tipe data:", tipe2)
}

type Person struct {
	name string
	age  int
}

func (p Person) Introduce(message string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", message, p.name, p.age)
}

func (p Person) ChangeName1() {
	p.name = "Danar"
}

func (p *Person) ChangeName2() {
	p.name = "Danar"
}

// Function

func sum(a, b int) int {
	return a + b
}

func calculate(a, b float64) (float64, float64) {
	multiple := a * b
	division := a / b
	return multiple, division
}

func sumAll(numbers ...int) int {
	total := 0
	for _, val := range numbers {
		total += val
	}
	return total
}
