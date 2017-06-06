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
	x := 100
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
	u := User{1, "ok", 12}
	Info(u)
	Set(&u)
	fmt.Println(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Typeof", t.Name(), t.Kind())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("can't set")
		return
	}

	v = v.Elem()
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("没找到")
		return
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("改变就是好事")
	}
}
