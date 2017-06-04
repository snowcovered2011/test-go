package main

import "fmt"

type A struct {
	Name string
}

func (a A) Print(name string) {
	fmt.Println(name)
}

func main() {
	a := A{Name: "jmxb"}
	a.Print(a.Name)
}
