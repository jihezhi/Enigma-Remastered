package main

import (
	"fmt"
)

func step(in *chan int, out *chan int) {
	x := <-*in
	*out <- x + 1
}

func main() {
	x := 0
	input := make(chan int)
	start := &input
	var end *chan int

	for i := 0; i < 10000; i++ {
		temp := make(chan int)
		end = &temp
		go step(start, end)
		start = end
	}
	input <- x
	result := <-*end
	fmt.Println(result)

}
