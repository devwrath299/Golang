package main

import "fmt"

func main() {

	// The main idea behind learning a pointer it gurantees that actual value is 
	//passed in function not the copy and also reflect changes in that variable.

	var ptr *int
	no := 34
	ptr = &no
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr=*ptr*2
	fmt.Print(*ptr)

}