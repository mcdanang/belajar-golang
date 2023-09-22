package main

import "fmt"

func main() {
	var word string = "selamat malam"

	var arr = make([]string, len(word))
	var dictionary = map[string]int{}
	for index, value := range word {
		fmt.Println(string(value))
		if dictionary[string(value)] == 0 {
			dictionary[string(value)] = 1
		} else {
			dictionary[string(value)]++
		}
		arr[index] = string(value)
	}
	fmt.Println(dictionary)
}
