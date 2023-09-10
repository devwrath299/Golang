package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	for ind, val := range arr {
		fmt.Println("at index", ind, " value is", val)
	}
}