package main

import "fmt"

func main() {

	var i = 1

	switch i {

	case 1:
		{
			fmt.Println("i is 1")
		}
	case 2:
		{
			fmt.Println("i is 2")
		}
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println(i, " is an integer")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("i is of a different type")
		}
	}

	whatAmI(1)
	whatAmI("hello")
	whatAmI(true)
}
