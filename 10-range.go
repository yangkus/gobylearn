package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5}

	for index, value := range num {
		fmt.Println("Index:", index, "Value:", value)
	}

}
