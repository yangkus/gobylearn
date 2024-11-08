package main

import "fmt"

func ptr(i int) {
	i = 0
}

func ptrr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println(i)

	ptr(i)

	prr := &i

	fmt.Println(i)

	ptrr(prr)

	fmt.Println(i)
}
