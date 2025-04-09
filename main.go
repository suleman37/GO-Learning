package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	now := time.Now()

	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileContent := string(file)

	fmt.Println("Total Number of Characters :", len(fileContent))
	words, line, spaces, vowels, punchar , digits , cons , special := combination(fileContent)
	fmt.Println("Total Number of words       :", words)
	fmt.Println("Total Number of Lines       :", line)
	fmt.Println("Total Number of Spaces      :", spaces-1)
	fmt.Println("Total Number of Vowels      :", vowels)
	fmt.Println("Total Number of Punctuations:", punchar)
	fmt.Println("Total Number of Digits   :", digits)
	fmt.Println("Total Number of Vowels     :", cons)
	fmt.Println("Total Number of Special Characters    :", special)
	timeTaken := time.Now().Sub(now)
	fmt.Println("Time taken in whole processing: ", timeTaken)
}

func combination(fileContent string) (int, int, int, int, int , int , int , int) {
	punchar := 0
	special :=0
	spaces := 0
	vowels := 0
	words := 0
	line := 0
	cons := 0
	digits := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == ' ' {
			words++
		}
		if fileContent[i] == '\n' {
			line++
		}
		if fileContent[i] == ' ' {
			spaces++

		}
		switch fileContent[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		switch fileContent[i] {
		case '.', ',', '\'', '(', ')', '!', '?', ';', ':', '-', '"':
			punchar++
		}
		switch fileContent[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digits++
		}
		if fileContent[i] != 'a' && fileContent[i] != 'e' && fileContent[i] != 'i' && fileContent[i] != 'o' && fileContent[i] != 'u' {
			cons++
		}
		if !(fileContent[i] >= 'a' && fileContent[i] <= 'z') && !(fileContent[i] >= 'A' && fileContent[i] <= 'Z') && !(fileContent[i] >= '0' && fileContent[i] <= '9') && fileContent[i] != ' ' {
			special++
		}
	}
	return words, line, spaces, vowels, punchar , digits ,cons , special
}