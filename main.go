package main

import (
	"fmt"
	"gotask/analyzer"
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
	lengthofData := len(fileContent);

	chunksize := lengthofData / 8

	fmt.Println(chunksize)
	chunks := make([]string,8);
	for i := 0; i < 8; i++ {
		startChunk := i * chunksize;
		endChunk := (i + 1) * chunksize;
		if i == 8-1 {
			endChunk = lengthofData
		}
		chunks[i] = fileContent[startChunk:endChunk]
	}


	fmt.Println("Total Number of Characters  :", len(fileContent))
	words, line, spaces, vowels, punchar, digits, cons, special := analyzer.Combination(fileContent)
	fmt.Println("Total Number of words       :", words)
	fmt.Println("Total Number of Lines       :", line)
	fmt.Println("Total Number of Spaces      :", spaces-1)
	fmt.Println("Total Number of Vowels      :", vowels)
	fmt.Println("Total Number of Punctuations:", punchar)
	fmt.Println("Total Number of Digits      :", digits)
	fmt.Println("Total Number of Vowels      :", cons)
	fmt.Println("Total Special Characters    :", special)
	timeTaken := time.Since(now)
	fmt.Println("Time taken in whole processing:", timeTaken)
}