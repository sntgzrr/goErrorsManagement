package main

import (
	"errors"
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (r raizError) Error() string {
	return fmt.Sprintf("Error matemático: %v, %v, %v", r.lat, r.long, r.err)
}

func main() {
	_, err := raiz(-10)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		return f, raizError{"50.2289 N", "99.4656 W", errors.New(fmt.Sprintf("Raíz cuadrada negativa del número %v", f))}
	}
	return f, nil
}
