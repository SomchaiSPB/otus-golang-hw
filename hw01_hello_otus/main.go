package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"

	fmt.Println(reverseStr(str))
}

func reverseStr(str string) string {
	return stringutil.Reverse(str)
}