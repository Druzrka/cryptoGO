package main

import (
	"fmt"
	"bufio"
	"os"
)
const (
	ALPHABETICAL_LENGTH = 26

	ASCII_OF_A = 65
	ASCII_OF_Z = 90

	ASCII_OF_a = 97
	ASCII_OF_z = 122
)

func checkShiftNumber(shiftNumber rune) bool {
	// если в руну приходит буква, то она становиться нулем
	if shiftNumber != 0 {
		return true
	}
	return false
}

func getText()string {
	fmt.Print("Enter your text: ")
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	return text.Text()
}

func getShiftNumber() byte {
	var number byte
	fmt.Print("Enter shift number: ")
	fmt.Scanf("%b", &number)
	if checkShiftNumber(rune(number)) {
		return number
	}
	fmt.Println("ERROR! You entered letter")
	return 0
}

func LetterIsBig(letter byte) bool{
	if letter >= ASCII_OF_A && letter <= ASCII_OF_Z {
		return true
	}
	return false
}

func LetterIsLittle(letter byte) bool{
	if letter >= ASCII_OF_a && letter <= ASCII_OF_z {
		return true
	}
	return false
}

func letterEncryption(letter, shiftNumber byte) string {
	if LetterIsBig(letter) {
		return string(ASCII_OF_A + (letter - ASCII_OF_A + shiftNumber) % ALPHABETICAL_LENGTH)
	} else if LetterIsLittle(letter) {
		return string(ASCII_OF_a + (letter - ASCII_OF_a + shiftNumber) % ALPHABETICAL_LENGTH)
	}
	return string(letter)
}

func textEncryption(shiftNumber byte, text string) string {
	length := len(text)
	encryptedText := ""
	for i := 0; i < length; i++ {
			encryptedText += letterEncryption(text[i], shiftNumber)
	}
	return encryptedText
}

func main() {
	text := getText()
	shiftNumber := getShiftNumber()
	fmt.Printf("%s",textEncryption(shiftNumber, text))
}
