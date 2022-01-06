// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gookit/color"
)

// Formula is
type Formula struct{}

// Run is
func (f Formula) Run(writer io.Writer) {
	cmd := "rit math sum --number_1=10 --number_2=15"
	err := RunFormula(cmd)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ReadOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum, _ := strconv.Atoi(data.Output["sum"])
	color.Blue.Printf("Result is: %v", sum*5)

	defer func() {
		if err := RemoveFile(); err != nil {
			fmt.Println(err)
			return
		}
	}()
}

/*
	MELHORIAS:
	1. Se eu interrompo o comando, o arquivo é gerado da mesma forma
	2. Se eu rodar 2 fórmulas ao mesmo tempo, o arquivo será sobrescrito (até onde isso é um problema)
	3. Se não for fórmula encadeada, mas um script chamando uma fórmula?
*/
