package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Printf("Say your name: ")
	var yourName string
	_, err := fmt.Scanln(&yourName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("said: %s", yourName)

}
