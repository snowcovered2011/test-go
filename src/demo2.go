package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "aa", 2: "bb", 3: "cc", 4: "dd"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
		fmt.Println(m2[v])
	}
	fmt.Println(m2)
}
