package main

import (
	_ "fmt"
)

func sum(s []int, c chan int) {
	for _, v := range s {
		c <- v
		println("存入数据: ", v)
	}
	println("所有数字都放进缓存了")
	close(c)
}

func main() {
	var a int = 30
	switch a {
	case 1:
		println(123)
	default:
		println(3456)
	}
}
