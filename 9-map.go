package main

import "fmt"

func main() {

	var m = make(map[string]int)

	m["apple"] = 5
	m["banana"] = 6
	m["pare"] = 4

	fmt.Println(m)

	fmt.Println(m["apple"])

	v1 := m["banana"]
	fmt.Println(v1)

	delete(m, "banana")
	fmt.Println(m)

	a, b := m["apple"]

	fmt.Println(a, b)

}
