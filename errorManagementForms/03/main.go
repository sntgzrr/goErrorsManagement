package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sb, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(sb))
}
