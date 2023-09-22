package main

import "fmt"

func main() {
	fruits1 := []string{"apple", "mango", "orange", "banana", "grape"}
	fmt.Printf("fruits1: %#v\n", fruits1)
	// fmt.Printf("%v\n", fruits1)

	var fruits2 = fruits1[1:4]
	fmt.Printf("fruits2: %#v\n", fruits2)

	var fruits3 = fruits1[len(fruits1)-2:]
	fmt.Printf("fruits3: %#v\n", fruits3)

	var fruits4 = append(fruits1[:3], "durian")
	fmt.Printf("fruits4: %#v\n", fruits4)

	var fruits5 = append([]string{"durian"}, fruits1[:3]...)
	fmt.Printf("fruits5: %#v\n", fruits5)

	var fruits6 = fruits1[2:4]
	fmt.Printf("fruits6: %#v\n", fruits6)

	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println(firstNumber)
	fmt.Println(&firstNumber)
	fmt.Println(*secondNumber)
	fmt.Println(secondNumber)

	var age int = 10
	fmt.Println(age)
	changeValue(&age)
	fmt.Println(age)
}

func changeValue(val *int) {
	*val = 20
}
