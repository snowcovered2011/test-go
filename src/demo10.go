package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}
func main() {
	u := User{1, "ok", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Typeof", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println("%6s: %v = %v", f.Name, f.Type, val)
	}
}
