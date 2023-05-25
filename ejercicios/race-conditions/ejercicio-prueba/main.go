package main

import (
	"fmt"
	"sync"
)

var counter int
var m sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				m.Lock()
				counter++
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", counter)
}
