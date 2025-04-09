package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileContent := string(file)

	fmt.Println("Total Number of Characters :", len(fileContent))
	// Add Number of Words Logic
	words := findWordCount(fileContent)
	fmt.Println("Total Number of words      :", words)

	// Add Number of Lines Logic
	lines := findlinesCount(fileContent)
	fmt.Println("Total Number of Lines      :", lines)

	// Add Number of Spaces Logic
	spaces := findspacesCount(fileContent)
	fmt.Println("Total Number of Spaces      :", spaces-1)

	// Add Number of vowels Logic
	vowels := findvowelsCount(fileContent)
	fmt.Println("Total Number of Vowels     :", vowels)

	// Add Number of Punctuations Logic
	punchar := findPunctuationsCount(fileContent)
	fmt.Println("Total Number of Punctuations:", punchar)

	// Add Number of Digits Logic
	digits := findDigitsCount(fileContent)
	fmt.Println("Total Number of Digits   :", digits)

	// Add Number of Digits Logic
	cons := findConsonentsCount(fileContent)
	fmt.Println("Total Number of Vowels     :", cons)

	// Add Number of Special Characters Logic
	special := findspecialCount(fileContent)
	fmt.Println("Total Number of Special Characters    :", special)
}

// Seperate Funcation for words count
func findWordCount(fileContent string) int {
	words := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == ' ' {
			words++
		}
	}

	return words
}

// Seperate Funcation for lines count
func findlinesCount(fileContent string) int {
	lines := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == '\n' {
			lines++
		}
	}
	return lines
}

// Seperate Funcation for Spaces count
func findspacesCount(fileContent string) int {
	spaces := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == ' ' {
			spaces++
		}
	}
	return spaces
}

// Seperate Funcation for Spaces count
func findvowelsCount(fileContent string) int {
	vowels := 0
	for i := 0; i < len(fileContent); i++ {
		switch fileContent[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	return vowels
}

// Seperate Funcation for Punctuations count
func findPunctuationsCount(fileContent string) int {
	punchar := 0
	for i := 0; i < len(fileContent); i++ {
		switch fileContent[i] {
		case '.', ',', '\'', '(', ')', '!', '?', ';', ':', '-', '"':
			punchar++
		}
	}
	return punchar
}

// Seperate Funcation for Digits count
func findDigitsCount(fileContent string) int {
	digits := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] >= '0' && fileContent[i] <= '9' {
			digits++
		}
	}
	return digits
}

// Seperate Funcation for Consonents count
func findConsonentsCount(fileContent string) int {
	cons := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 'a' && fileContent[i] != 'e' && fileContent[i] != 'i' && fileContent[i] != 'o' && fileContent[i] != 'u' {
			cons++
		}
	}
	return cons
}

// Seperate Funcation for Special Characters count
func findspecialCount(fileContent string) int {
	special := 0
	for i := 0; i < len(fileContent); i++ {
		if !(fileContent[i] >= 'a' && fileContent[i] <= 'z') && !(fileContent[i] >= 'A' && fileContent[i] <= 'Z') && !(fileContent[i] >= '0' && fileContent[i] <= '9') && fileContent[i] != ' ' {
			special++
		}
	}
	return special
}
