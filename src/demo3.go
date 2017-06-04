package main

import "fmt"

func main() {
	var a = 1
	Test(&a)
	fmt.Println("11111")
	fmt.Println(a)
}
func Test(a *int) {
	fmt.Println(*a)
	*a = 3
}
