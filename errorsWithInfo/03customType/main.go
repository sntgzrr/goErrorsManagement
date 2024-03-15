package main

import (
	"fmt"
	"log"
)

type ubicationError struct {
	lat  string
	long string
	err  error
}

func (n ubicationError) Error() string {
	return fmt.Sprintf("Un error matemático ocurrió en: %v, %v, %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("Error mat: Raíz cuadrada de número negativo %v", f)
		return 0, ubicationError{"50.005 N", "99.5643 W", err}
	}
	return 7, nil
}
