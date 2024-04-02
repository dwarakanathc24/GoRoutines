package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int = 0

	var wg sync.WaitGroup
	var gr int = 100
	wg.Add(gr)

	go func() {
		for i := 0; i < gr; i++ {
			c := counter
			runtime.Gosched()
			c++
			counter = c
			wg.Done()
			fmt.Println("GoRoutine", runtime.NumGoroutine())
		}
	}()
	wg.Wait()
	fmt.Println("counter:", counter)

}
