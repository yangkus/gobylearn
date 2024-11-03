package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	fmt.Println(reflect.TypeOf(i)) // int

	var whatAmI = func(i interface{}) {

		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Println("I am of unknown type")
		}
	}

	whatAmI(1)
	whatAmI("hello")
	whatAmI(true)
	whatAmI(nil)
}
