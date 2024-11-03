package main

import (
	"fmt"
)

func main() {

	var i = 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("infinite loop")
		break
	}

}
