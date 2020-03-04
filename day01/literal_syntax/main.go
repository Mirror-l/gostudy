package main

import "fmt"

func main() {
	a := 0b0100
	fmt.Println(a)        // 二进制 1 * 2^2 = 4
	fmt.Printf("%d\n", a) // 以十进制显示
	fmt.Printf("%b\n", a) // 以二进制显示
	fmt.Printf("%o\n", a) // 以八进制显示

	b := 0100
	fmt.Println(b)        // 八进制 1 * 8^2 = 64
	fmt.Printf("%x\n", b) // 以十六进制显示

	c := 0x100
	fmt.Println(c) // 十六进制 1 * 16^2 = 256

	d := 123_456
	fmt.Println(d) // 连接符

	e := 15
	fmt.Printf("%x \n", e)
	fmt.Printf("%X \n", e)

}
