package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	myString := "❗hello"

	for index, char := range myString {
		fmt.Printf("Index(byte offset): %d\tCharacter: %c\tCode Point: %U\n",
			index, char, char)
	}

	byteLength := len(myString)
	runeLength := utf8.RuneCountInString(myString)
	fmt.Printf("Byte length: %d\n", byteLength)
	fmt.Printf("Rune count : %d\n", runeLength)

	// truy cập theo byte
	fmt.Printf("myString[0] (byte): %v\n", myString[0])

	// chuyển sang []rune để truy cập theo rune index
	runes := []rune(myString)
	fmt.Printf("runes[0]: %c (Code Point: %U)\n", runes[0], runes[0])
}
