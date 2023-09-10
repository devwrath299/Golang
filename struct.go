package main

import "fmt"

func main() {
	det := user{"Ram", 102}
	fmt.Println(det.get())
}

type user struct {
	Name string
	Age  int
}

func (u user ) get() string{
	return u.Name
}