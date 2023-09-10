package main

import "fmt"

func main() {

	maps := make(map[string]string)
	maps["M"] = "maths"
	maps["sc"] = "science"
	maps["his"] = "history"
	fmt.Println(maps)

	// to deleete particular coloumn
	delete(maps,"M")
	fmt.Println(maps)

	// for loop over map
	for k,v := range maps{
		fmt.Printf("for key %v,the value is %v \n",k,v)
	}

}