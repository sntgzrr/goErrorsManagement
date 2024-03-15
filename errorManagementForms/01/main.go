package main

import "fmt"

func main() {
	var name, lastname, age string
	fmt.Print("Nombre:")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Print("Apellido:")
	_, err = fmt.Scan(&lastname)
	if err != nil {
		panic(err)
	}

	fmt.Print("Edad:")
	_, err = fmt.Scan(&age)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tu nombre es", name, lastname, "y tienes una edad de", age, "años.")
}
