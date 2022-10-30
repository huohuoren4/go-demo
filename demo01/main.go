package main

import (
	_ "errors"
	"fmt"
)

func Divide(dividee int, divider int) (int, error) {
	if divider == 0 {
		return 0, fmt.Errorf("%d 和 %d 相除报错", dividee, divider)
	}
	return dividee / divider, fmt.Errorf("")
}

func main() {
	fmt.Println(Divide(10, 0))
	fmt.Println(Divide(10, 50))

}
