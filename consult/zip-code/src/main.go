package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	zipInput := os.Getenv("ZIP")

	formula.Input{
		Zip: zipInput,
	}.Run()
}
