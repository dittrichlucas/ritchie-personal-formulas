package formula

import (
	"fmt"
	"io"
	"strconv"
)

type Formula struct {
	Number1 string
	Number2 string
}

func (h Formula) Run(writer io.Writer) {
	n1, _ := strconv.Atoi(h.Number1)
	n2, _ := strconv.Atoi(h.Number2)

	if err := RunFormula("rit -v"); err != nil {
		fmt.Println(err)
	}
	sum := n1 + n2
	key := "math"
	value := map[string]interface{}{"inputs": map[string]int{"n1": n1, "n2": n2}, "sum": sum}

	SetOutput(key, value)

	RemoveOutput("sum")
}
