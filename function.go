package main

import "fmt"

func main() {
	val, ans := add(3, 4)
	fmt.Println(val, ans)
}

func add(a int, b int) (int, string) {
	return a + b, "Hello"
}