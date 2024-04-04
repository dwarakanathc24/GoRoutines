package main

import "fmt"

func main() {

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	m := map[string]string{
		"name": "ram", "email": "siva@gmail.com",
	}

	fmt.Println(m)

	go send(odd, even, quit)
	receive(odd, even, quit)

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	q <- 0

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the channle e", v)
		case v := <-o:
			fmt.Println("from the channle o", v)
		case v := <-q:
			fmt.Println("from the channle q", v)
			return

		}
	}

}
