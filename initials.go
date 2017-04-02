package main

import(
	"fmt"
	"strings"
)

func initials(words string) {

	length := len(words)
	const first = 0
	if words[first] != ' ' {
		fmt.Printf("%c", words[0])
	}

	for index := range words {
		if words[index] == ' ' && index + 1 != length && words[index + 1] != ' '  {
			fmt.Printf("%c", words[index + 1])
		}
	}
}

func main() {
	var words = "   hello      my     friend   "
	words = strings.ToUpper(words)
	initials(words)
}
