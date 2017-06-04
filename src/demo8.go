package main

import "fmt"

type A int

func main() {
	var a A
	a = 1
	a.Add(100)
	fmt.Println(a)
}

func (a *A) Add(b int) {
	*a += A(b)
}
