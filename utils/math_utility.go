package utils

import (
	"bytes"
)

func GetArrayByAddingOne() []int {

	var arrval = []int{2, 3, 4, 5}

	for i := 0; i < len(arrval); i++ {
		arrval[i] = arrval[i] + 1
	}
	return arrval
}

func GetEvenNumber() []int {

	var arrval = []int{2, 3, 4, 5}
	var evenNum []int
	for i := 0; i < len(arrval); i++ {
		if arrval[i]%2 == 0 {
			evenNum = append(evenNum, arrval[i])
		}
	}
	return evenNum
}

func GetOddNumber() []int {

	var arrval = []int{2, 3, 4, 5}
	var evenNum []int
	for i := 0; i < len(arrval); i++ {
		if arrval[i]%2 != 0 {
			evenNum = append(evenNum, arrval[i])
		}
	}
	return evenNum
}

func ConcatStrings(str ...string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(str); i++ {

		buffer.WriteString(str[i])

	}
	return buffer.String()
}
