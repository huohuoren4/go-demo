package main

import (
	"bufio"
	"fmt"
	"os"
)

type Phone interface {
	call() int
}
type SumsungPhone struct {
	phoneType string
}

func (sphone *SumsungPhone) call() int {
	fmt.Println("This is a call from", sphone.phoneType)
	return 2
}

func main() {
	bufio.NewReader(os.Stdin)
}
