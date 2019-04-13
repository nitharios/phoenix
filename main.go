package main

import (
	"fmt"
	"os"
)

func main() {
	ENV := os.Getenv("ENV")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	s := New(HOST)

	if ENV != "PROD" {
		fmt.Println("Starting test server on ", PORT)
		s.Logger.Fatal(s.Start(PORT))

	} else {
		fmt.Println("Starting production server on ", PORT)
		s.Logger.Fatal(s.Start(PORT))
	}
}
