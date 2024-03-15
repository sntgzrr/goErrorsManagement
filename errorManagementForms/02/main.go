package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Se crea un archivo .txt
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Al finalizar el main, se cierra el file creado.
	defer f.Close()

	//Se crea r para guardar un string Reader.
	r := strings.NewReader("Santiago Lozano")

	//Escribo el r dentro del file .txt f
	io.Copy(f, r)
}
