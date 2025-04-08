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
	a :=string(file)

	fmt.Println("Total Number of Characters :",len(a))

	// Add Number of Words Logic
    words := 0
	for i := 0; i < len(a); i++ {
		if a[i] == ' '{
			words++
		}
	}
	fmt.Println("Total Number of words      :",words)
}