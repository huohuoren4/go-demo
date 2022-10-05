package main

import (
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
	var phone Phone
	fmt.Printf("%T", phone)
}
