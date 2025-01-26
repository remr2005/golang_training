package main

import (
	"fmt"
)

func F_Hello(s int, n chan<- int) {
	fmt.Println("Hello")
	n <- s * s
}

func F_World(s int, n <-chan int) {
	fmt.Println("World")
	fmt.Println(<-n)
}

func main() {
	n := make(chan int)
	go F_Hello(2, n)
	go F_World(3, n)
	<-n
}
