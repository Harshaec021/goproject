package main

import (
	"fmt"

	"github.com/Harshaec021/goproject/morestrings"

	"github.com/Harshaec021/goproject/utils"

	r "math"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!GooG ,olleH"))
	fmt.Println(utils.GetGreeting("Harsha"))
	fmt.Println(utils.GetArrayByAddingOne())
	fmt.Println(utils.GetEvenNumber())
	fmt.Println(utils.GetOddNumber())
	fmt.Println(utils.ConcatStrings("Hello", "World"))
	fmt.Println(r.Sqrt(4))
}
