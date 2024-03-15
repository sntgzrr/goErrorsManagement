package main

import "fmt"

type erroPer struct {
	info string
}

func (e erroPer) Error() string {
	return "Un error ha ocurrido."
}

func main() {
	err := erroPer{
		info: "Info",
	}
	foo(err)
}

func foo(e error) {
	fmt.Println(e.Error(), e.(erroPer).info)
}
