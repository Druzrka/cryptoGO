package main

import (
	"fmt"
)

func isBig(letter byte) bool{
	if letter > 64 && letter < 91 {
		return true
	}
	return false
}

func isLittle(letter byte) bool{
	if letter > 96 && letter < 123 {
		return true
	}
	return false
}

func cipher(shiftNumber byte, text string) {
	length := len(text)
	for i := 0; i < length; i++ {
		if isBig(text[i]){
			fmt.Printf("%c",'A' + (text[i] - 'A' + shiftNumber) % 26)
		} else if isLittle(text[i]) {
			fmt.Printf("%c",'a' + (text[i] - 'a' + shiftNumber) % 26)
		} else {
			fmt.Printf("%c", text[i])
		}
	}
}

func main() {
	var shiftNumber byte = 4
	var text string = "HELLO my dear friend!"
	cipher(shiftNumber, text)
}
