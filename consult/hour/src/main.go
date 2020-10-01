package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	firstEntry := os.Getenv("FIRST_ENTRY")
	firstExit := os.Getenv("FIRST_EXIT")
	secondEntry := os.Getenv("SECOND_ENTRY")
	secondExit := os.Getenv("SECOND_EXIT")

	formula.Formula{
		FirstEntry:  firstEntry,
		FirstExit:   firstExit,
		SecondEntry: secondEntry,
		SecondExit:  secondExit,
	}.Run(os.Stdout)
}
