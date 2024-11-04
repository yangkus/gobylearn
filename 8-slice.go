package main

import "fmt"

func main() {
	var s = make([]int, 5)

	s = append(s, 2)

	fmt.Println("Length of the slice:", len(s))
	fmt.Println(s)

	l := s[2:4]
	fmt.Println("New slice:", l)

	l = s[2:]
	fmt.Println("New slice:", l)
}
