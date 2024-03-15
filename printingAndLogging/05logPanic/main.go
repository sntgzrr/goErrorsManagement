package main

import "os"

func main() {
	_, err := os.Open("no_found.txt")
	if err != nil {
		panic(err)
	}
}
