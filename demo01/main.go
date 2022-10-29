package main

import (
	"demo01/test"
	"fmt"
)

type Phone interface {
	call() int
}

type ApplePhone struct {
	phoneType string
}

type SumsungPhone struct {
	phoneType string
}

func (iphone ApplePhone) call() int {
	fmt.Println("This is a call from", iphone.phoneType)
	return 1
}

func (sphone SumsungPhone) call() int {
	fmt.Println("This is a call from", sphone.phoneType)
	return 2
}

func main() {
	phone := new(SumsungPhone)
	fmt.Println(phone)
	fmt.Println("hello, world")
	fmt.Printf("test.Add(3, 4): %v\n", test.Add(3, 4))
}
