package main

import (
	"fmt"
)

/*
 * 不能直接运行，会报undefine错误
 * 在控制台中输入go run XXX.go main.go
 */
func main() {
	strings := make([]int, 0, 10)
	strings = append(strings, 4)
	strings = append(strings, 3)
	strings = append(strings, 11)
	strings = append(strings, 8)

	InsertSort(strings)
	fmt.Println(strings)
}
