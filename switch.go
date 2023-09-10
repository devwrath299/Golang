package main

import (
	"fmt"
	"math/rand"

)

func main() {
	value:=rand.Intn(6)

	switch value{
	case 1:fmt.Println("1 here")
	case 2:fmt.Println("2 here")
	case 3:fmt.Println("3 here")
	fallthrough
	case 4:fmt.Println("4 here")
	fallthrough
	case 5:fmt.Println("5 here")
	default:fmt.Println("6 here")
	}

}
