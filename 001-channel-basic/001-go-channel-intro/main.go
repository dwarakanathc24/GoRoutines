package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(3)
	a := 0
	c := make(chan int)
	count := 100

	for i := 0; i < count; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			c <- 75
		}()
	}

	for i := 0; i < count; i++ {
		fmt.Println("the channel is", <-c)
	}

	groutine1(&a)
	groutine2(&a)
	groutine3(&a)
	wg.Wait()

	fmt.Println("the final value of a is:", a)
}

func groutine1(a *int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			*a++
			mutex.Unlock()
		}()
	}
}

func groutine2(a *int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			*a++
			mutex.Unlock()
		}()
	}
}

func groutine3(a *int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			*a++
			mutex.Unlock()
		}()
	}
}
