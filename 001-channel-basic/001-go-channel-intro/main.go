package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 75

	fmt.Println(" the channel is ", <-c)
	fmt.Println(" the channel is ", c)
	fmt.Printf(" the c type is %T ", c)

}
