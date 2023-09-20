package main

import (
	"fmt"
	"strconv"
)

const (
	pi = 3.14
)

const (
	statusNew = iota + 1
	statusOld
)

func main() {
	// fmt.Println("Hello world!")/

	var title string = "Golang for everyone"
	var member int
	member = 212

	fmt.Println(title)
	fmt.Println(member)

	title2 := "Golang for anyone"
	member2 := 717

	fmt.Printf("%T, %v \n", title2, title2)
	fmt.Printf("%T, %v \n", member2, member2)

	var firstName, surName, lastName string = "Muhamad", "Danang", "Priambodo"
	age1, age2, _ := 20, "21", 22

	fmt.Println(firstName, surName, lastName)
	fmt.Println(age1, age2)
	fmt.Printf("Hello my name is %s and I'm %d years old\n", surName, age1)

	decimalNumber := 3.752
	fmt.Printf("Decimal number: %f \n", decimalNumber)
	fmt.Printf("Decimal number: %.2f \n", decimalNumber)

	// const pi = 3.14
	fmt.Println(pi)

	fmt.Println(statusNew)
	fmt.Println(statusOld)

	stringNumber := "1234"
	num, _ := strconv.Atoi(stringNumber)
	fmt.Println(num)
	fmt.Printf("Tipe data num %T\n", num)
}
