package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("NUMBER_1")
	input2 := os.Getenv("NUMBER_2")

	formula.Formula{
		Number1: input1,
		Number2: input2,
	}.Run(os.Stdout)
}
