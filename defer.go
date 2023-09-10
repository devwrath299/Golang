package main

import "fmt"

func main() {
	defer fmt.Println("w")
	defer fmt.Println("r")

	// defer statement runs at last
	// it follow Lifo pocess

}