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

	// 判断是否包含
	fmt.Println(strings.Contains("abc123", "123"))

	// 前缀判断
	fmt.Println(strings.HasPrefix("incompatible", "in"))

	// 后缀判断
	fmt.Println(strings.HasSuffix("incompatible", "able"))

	// 字串第一次出现的位置
	fmt.Println(strings.Index("i have a computer", "have a"))

	// 字串最后一次出现的位置
	fmt.Println(strings.LastIndex("hahaha", "h"))

	// 使用指定间隔字符合并字符串
	array := []string{"abc", "is", "123"}
	fmt.Println(strings.Join(array, ","))

}
