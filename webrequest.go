package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.indiamart.com/"

func main() {
	
	response,er:= http.Get(url)
	if er!=nil{
		panic(er)
	}

	data,er:=ioutil.ReadAll(response.Body)

	cont:=string(data)
	fmt.Println(cont)


}