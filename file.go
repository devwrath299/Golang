package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	con := "Hello I am new file"

	file, err := os.Create("./cr.txt")
    
	if err!=nil{
		panic(err)
	}

	c,er :=io.WriteString(file,con)
	if er!=nil{
		panic(er)
	}
	fmt.Println(c)
	defer file.Close()

}