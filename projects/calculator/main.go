package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		select {
		case x := <-firstChan:
			out <- x * x
		case x := <-secondChan:
			out <- x * 3
		case <-stopChan:
			return
		}
	}()
	return out
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	stopchan := make(chan struct{})
	resalt := calculator(channel1, channel2, stopchan)
	//channel1 <- 4
	channel2 <- 2
	fmt.Println(<-resalt)
}
