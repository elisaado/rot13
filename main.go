// ROT13 by Eli Saado
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Returns whether a string is uppercase or not
func isUpper(str string) bool {
	return strings.ToUpper(str) == str
}

// Rotates a char by 13 positions in the alphabet
func rot13(char string, alphabets [][]byte) string {
	var alphabet []byte

	if isUpper(string(char)) {
		alphabet = alphabets[0]
	} else {
		alphabet = alphabets[1]
	}

	index := bytes.IndexByte(alphabet, []byte(char)[0])
	if index < 0 {
		return char
	} else if index >= 13 {
		index -= 13
	} else {
		index += 13
	}

	return string(alphabet[index])
}

func main() {
	var CharsRot13 string
	for _, char := range strings.Split(strings.Join(os.Args[1:], " "), "") {
		CharsRot13 += rot13(char, [][]byte{[]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}, []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}})
	}

	fmt.Println(CharsRot13)
}
