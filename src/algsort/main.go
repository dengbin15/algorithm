package main

import (
	"fmt"
)

func main() {
	strings := make([]int, 0, 10)
	strings = append(strings, 4)
	strings = append(strings, 3)
	strings = append(strings, 11)
	strings = append(strings, 8)

	InsertSort(strings)
	fmt.Println(strings)
}
