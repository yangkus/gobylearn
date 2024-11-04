package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("Initial array:", a)

	a[0] = 100

	fmt.Println(a[0])

	var b = [5]int{0, 1, 2, 3, 4}

	fmt.Println("Array after initialization:", b)

}
