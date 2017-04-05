package main

import (
	"fmt"
	"strings"
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

func getText() string {
	fmt.Print("Enter your text: ")
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	return text.Text()
}

func getShiftWord() string {
	fmt.Print("Enter the key: ")
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	return text.Text()
}

func LetterIsBig(letter byte) bool {
	if letter >= ASCII_OF_A && letter <= ASCII_OF_Z {
		return true
	}
	return false
}

func LetterIsLittle(letter byte) bool {
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

func textEncryption(shiftNumber []byte, text string) string {
	lengthOfText := len(text)
	lengthOfShiftNumber := len (shiftNumber)
	encryptedText := ""
	for i := 0; i < lengthOfText; i++ {
		encryptedText += letterEncryption(text[i], shiftNumber[i % lengthOfShiftNumber])
	}
	return encryptedText
}

func generateShiftNumbers(shiftWord string) []byte {
	shiftWord = strings.ToLower(shiftWord)
	length := len(shiftWord)
	shiftNumbers := make([]byte, length)
	for i := 0; i < length; i++ {
		shiftNumbers[i] = shiftWord[i] - ASCII_OF_a
	}
	return shiftNumbers
}

func main() {
	shiftWord := getShiftWord()
	text  := getText()
	shiftNumbers := generateShiftNumbers(shiftWord)
	fmt.Println(textEncryption(shiftNumbers,text))
}
