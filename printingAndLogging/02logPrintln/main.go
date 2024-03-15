package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no_found.txt")
	if err != nil {
		log.Println("An error has occurred:", err)
	}
}
