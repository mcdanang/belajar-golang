package main

import (
	"fmt"
	"sync"
)

type printData interface {
	print(index int, wg *sync.WaitGroup)
}

type data struct {
	text []string
}

func (d data) print(index int, wg *sync.WaitGroup) {
	fmt.Printf("%s %v\n", d.text, index)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(8)
	var coba printData = data{[]string{"coba1", "coba2", "coba3"}}
	var bisa printData = data{[]string{"bisa1", "bisa2", "bisa3"}}
	for i := 1; i <= 4; i++ {
		go coba.print(i, &wg)
		go bisa.print(i, &wg)
	}
	wg.Wait()
}
