package main

import (
	"fmt"
	"sync"
)

var Counter int
var m sync.Mutex

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go incremental(&wg)
	go incremental(&wg)
	wg.Wait()

	fmt.Printf("La suma final del contador es: %d", Counter)
}

func incremental(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		m.Lock()
		Counter++
		m.Unlock()
	}
	wg.Done()
}
