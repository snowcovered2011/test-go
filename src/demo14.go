package main

import (
	"fmt"
)

func main() {
	c1, c2, o := make(chan int), make(chan string), make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)

			}
		}
	}()
	c1 <- 1
	c2 <- "ok"
	c1 <- 2
	c2 <- "hello"
	close(c1)
	close(c2)
	<-o
}
