package main

import "fmt"

func main() {
	var arr []int
	ar :=[]int{}
	arr=append(arr, 3, 4, 5)
	ar=append(arr[:2])
	fmt.Println(ar)

	// how to remove index
	subj := []string{"maths","science","history","Geography"}
	subj=append(subj[:2],subj[3:]...)
	fmt.Println(subj)




}