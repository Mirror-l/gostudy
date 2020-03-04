package main

import (
	"fmt"
	"strings"
)

func main() {

	// 字符串拼接
	fmt.Println("hello " + "world")

	fmt.Println(fmt.Sprintf("%s %s", "hello", "world"))

	// 字符串分割
	fmt.Println(strings.Split("1,2,3,", ","))

}
