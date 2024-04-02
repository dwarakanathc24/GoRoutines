package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(" speak() is calling")
	fmt.Println("person name is ", p.name)
}
func saySomething(h human) {
	fmt.Println(" this is the saysomething() calling")
	h.speak()
}

func main() {
	p := person{
		name: "abhay",
	}
	saySomething(&p)
}
