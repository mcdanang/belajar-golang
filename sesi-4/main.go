package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Goroutine
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printName("Danang", &wg)
	// go printName("Danar", &wg)
	// // time.Sleep(time.Millisecond * 100)
	// wg.Wait()

	// Defer example
	// defer fmt.Println("Haiii")
	// fmt.Println("Halooo")

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }

	// Exit example
	// defer fmt.Println("Invoke with defer")
	// fmt.Println("teees")
	// os.Exit(0)

	// Channel
	// s := []int{7, 2, 8, -9, 4, 9}
	// c := make(chan int)

	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)

	// x, y := <-c, <-c

	// fmt.Println(x, y, x+y)

	// Error & Panic
	// var number int
	// var err error

	// number, err = strconv.Atoi("12345")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// var password string = "test"
	// valid, err := validPassword(password)
	// fmt.Println(valid)
	// fmt.Println(err)
	// panic(err)
}

func printName(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	time.Sleep(time.Millisecond * 100)
	wg.Done()
}

func sum(x []int, c chan int) {
	sum := 0
	for _, v := range x {
		sum += v
	}

	c <- sum
}

func validPassword(password string) (string, error) {
	if len(password) < 5 {
		return "", errors.New(("Password has to have more than 4 characters"))
	}

	return "Valid password", nil
}
