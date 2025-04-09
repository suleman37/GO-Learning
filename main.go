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
	a := string(file)

	fmt.Println("Total Number of Characters :", len(a))

	// Add Number of Words Logic
	words := 0
	for i := 0; i < len(a); i++ {
		if a[i] == ' ' {
			words++
		}
	}
	fmt.Println("Total Number of words      :", words)

	// Add Number of Lines Logic
	lines := 0
	for i := 0; i < len(a); i++ {
		if a[i] == '\n' {
			lines++
		}
	}
	fmt.Println("Total Number of Lines      :", lines)

	// Add Number of Spaces Logic
	spaces := 0
	for i := 0; i < len(a); i++ {
		if a[i] == ' ' {
			spaces++
		}
	}
	fmt.Println("Total Number of Spaces     :", spaces-1)

	// Add Number of vowels Logic
	vowels := 0
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	fmt.Println("Total Number of Vowels     :", vowels)

	// Add Number of Punctuations Logic
	punchar := 0
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case '.', ',', '\'', '(', ')', '!', '?', ';', ':', '-', '"':
			punchar++
		}
	}
	fmt.Println("Total Number of Punctuations:", punchar)

	// Add Number of Digits Logic
	digits := 0
	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			digits++
		}
	}
	fmt.Println("Total Number of Digits   :", digits)

	// Add Number of consonents Logic
	cons := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 'a' && a[i] != 'e' && a[i] != 'i' && a[i] != 'o' && a[i] != 'u' {
			cons++
		}
	}
	fmt.Println("Total Number of Vowels     :", cons)


	// Add Number of Special Characters Logic
	special := 0
	for i := 0; i < len(a); i++ {
			if !(a[i] >= 'a' && a[i] <= 'z') && !(a[i] >= 'A' && a[i] <= 'Z') && !(a[i] >= '0' && a[i] <= '9') && a[i] != ' ' {
				special++
			}
	}
	fmt.Println("Total Number of Special Characters    :", special)
}