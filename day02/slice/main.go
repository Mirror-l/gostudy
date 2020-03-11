package main

import "fmt"

func main() {

	// 声明一个切片
	var a []int

	fmt.Println(a) // []

	a = append(a, 1, 2, 3, 4, 5)

	fmt.Println(a)

	a = append(a[:1], a[2:]...) // ... 打散切片
	fmt.Println(a)

	// 声明并初始化一个切片

	var b = []int{22, 33, 44}

	fmt.Println(b)

	// 从数组初始化切片

	var c = [4]int{1, 2, 3, 45}

	d := c[1:2] // 左闭右开

	fmt.Println(d)
}
