package main

import "fmt"

var ff = func(nums ...int) {

	fmt.Println(nums)

}

func main() {

	ff(1, 2, 3, 4, 5)

	num1 := []int{1, 2, 3, 4, 5}

	ff(num1...)
}
