package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Running the uuid lib")

	uuidStr := uuid.NewString()

	fmt.Printf("Generated: %s", uuidStr)

}
