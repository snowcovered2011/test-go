package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Name string
	Sex  string
}

type Teacher struct {
	Person
	Sex string
}

func main() {
	a := Person{
		Name: "yanlihui",
		Age:  26,
	}
	b := Student{Sex: "男", Name: "sfsda", Person: Person{Name: "jmxb", Age: 23}}
	c := Teacher{Sex: "男", Person: Person{Name: "jmxb", Age: 23}}
	b.Name = "snow"
	fmt.Println(a)
	Test(&a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func Test(per *Person) {
	per.Name = "jmxb"
	per.Age = 27
	fmt.Println(per)
}
