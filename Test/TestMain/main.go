package main

import (
	"fmt"

	"example.com/Test/lib"
)

func main() {
	var storeval int32
	var values []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go lib.PrintArr(values)
	fmt.Scanln(&storeval)
}
