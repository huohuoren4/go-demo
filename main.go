package main

import "fmt"

func main() {
	var i int = 123
	var str string = "123"
	var f float32 = 123.5646
	var p *int = &i
	fmt.Println(i, str)
	fmt.Printf("hello, %T, %.2f, %d", str, f, *p)
}
