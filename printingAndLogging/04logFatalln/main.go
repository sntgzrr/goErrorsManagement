package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no_found.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, the defer functions don't run.")
}
