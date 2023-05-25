package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Clear the screen by printing \x0c.

	wg.Add(1)
	go bar(&wg)
	into()
}

func bar(wg *sync.WaitGroup) {
	const col = 30

	bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
		if i == col-1 {
			i = 0
		}
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
	wg.Done()
}

func into() {
	fmt.Print("Hi")
}
