package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Println("控制语句相关:")

	// if 语句
	fmt.Println("if语句:")
	fmt.Println("请输入a, b:")
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Printf("当前 a:%d > b:%d ", a, b)
	} else if b > a {
		fmt.Printf("当前 a:%d < b:%d ", a, b)
	} else {
		fmt.Printf("当前 a:%d = b:%d ", a, b)
	}

}
