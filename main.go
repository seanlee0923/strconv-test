package main

import (
	"fmt"
	"strconv"
)

func BinaryStrToInt() {
	binaryStrings := [3]string{"01000111", "01001111", "01011111"}

	for _, binaryString := range binaryStrings {
		if val, err := strconv.ParseInt(binaryString, 2, 64); err == nil {
			fmt.Println("str: ", binaryString)
			fmt.Println("i: ", val)
			runes := rune(val)
			fmt.Println("rune: ", string(runes))
		}
	}

}
