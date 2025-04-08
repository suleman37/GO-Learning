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

	fmt.Println(len(a))

	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
}