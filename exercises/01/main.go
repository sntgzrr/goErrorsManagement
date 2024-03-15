package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name     string
	Lastname string
	Phrases  []string
}

func main() {
	f, err := os.Create("json.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p1 := person{
		Name:     "Santiago",
		Lastname: "Lozano",
		Phrases:  []string{"Lmao", "XD", "LOL"},
	}

	sb, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(sb)
	if err != nil {
		panic(err)
	}

	fmt.Println("See your json.txt file.")
}
