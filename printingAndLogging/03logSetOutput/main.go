package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
	defer f.Close()
	//Archivo en donde se escribirán los mensajes.
	log.SetOutput(f)

	f2, err := os.Open("no_found.txt")
	if err != nil {
		log.Println("An error has occurred:", err)
	}
	defer f2.Close()

	fmt.Println("Check log file in your directory.")
}
