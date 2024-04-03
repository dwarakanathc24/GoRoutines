package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println(" go routines 1")
		wg.Done()
	}()

	go func() {
		fmt.Println(" go routines 2")
		wg.Done()
	}()

	go func() {
		fmt.Println(" go routines 3")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("exit of the program")

}
