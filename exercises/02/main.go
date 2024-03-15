package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	Name     string
	Lastname string
	Phrases  []string
}

func main() {
	p1 := person{
		Name:     "Santiago",
		Lastname: "Lozano",
		Phrases:  []string{"Lmao", "XD", "LOL"},
	}

	sb, err := aJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sb))
}

func aJSON(a interface{}) ([]byte, error) {
	sb, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("Can't convert to JSON this value: %v", a))
	}
	return sb, nil
}
