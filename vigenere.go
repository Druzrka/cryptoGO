package main

import (
	"fmt"
	"strings"
)

func isBig(letter byte) bool{
	if letter > 64 && letter < 91 {
		return true
	}
	return false
}

func isLittle(letter byte) bool {
	if letter > 96 && letter < 123 {
		return true
	}
	return false
}

func cipher(shiftNumber []byte, text string) {
	lengthOfText := len(text)
	for i := 0; i < lengthOfText; i++ {
		if isBig(text[i]){
			fmt.Printf("%c",'A' + (text[i] - 'A' + shiftNumber[i % len(shiftNumber)]) % 26)
		} else if isLittle(text[i]) {
			fmt.Printf("%c",'a' + (text[i] - 'a' + shiftNumber[i % len(shiftNumber)]) % 26)
		} else {
			fmt.Printf("%c", text[i])
		}
	}
}

func shiftNumbers(shiftWord string) []byte {
	shiftWord = strings.ToLower(shiftWord)
	length := len(shiftWord)
	shiftNumbers := make([]byte, length)
	for i := 0; i < length; i++ {
		shiftNumbers[i] = shiftWord[i] - 'a'
	}
	return shiftNumbers
}

func main() {
	var shiftWord string = "word"
	var text string = "HELLO my dear friend!"
	var shiftNumbers []byte = shiftNumbers(shiftWord)
	cipher(shiftNumbers, text)
}
