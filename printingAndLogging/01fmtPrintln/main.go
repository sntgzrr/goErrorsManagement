package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no_found.txt")
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
}
