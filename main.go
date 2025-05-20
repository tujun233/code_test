package main

import "fmt"

type Person struct {
	Age int64 `json:"age,omitempty"`
	A   *User `json:"a"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	var p []int
	fmt.Println(p)
	change(p)
	fmt.Println(p)
	fmt.Println(345)
	fmt.Println(1234)
	fmt.Println("sdfsdf")
	fmt.Println("sdfsdf123123")
	fmt.Println("qwe")
}

func change(p []int) {
	p = append(p, 1)
}
