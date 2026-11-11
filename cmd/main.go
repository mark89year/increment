package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("specify number for increment: inc 42")
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		panic("number should be a number")
	}
	fmt.Println(n + 1)
}
