package main

func values() (int, string) {

	return 10, "hello"
}

func main() {
	a, b := values()
	println(a, b)

	_, c := values()
	println(c)
}
