// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"io"
	"os/exec"

	"github.com/gookit/color"
)

// Formula is
type Formula struct{}

/*
	TODO:
		1. Fazer a requisição para a URL
		2. Montar a string que contém a versão do Go que eu pretendo baixar
		3. Verificar se o download é realizado
		4. Realizar o update da versão
		5. Checar se a versão foi atualizada
*/

// Run is
func (f Formula) Run(writer io.Writer) {
	goVersion := exec.Command("go", "version")
	out, _ := goVersion.CombinedOutput()
	// https: //golang.org/dl/go1.14.12.linux-amd64.tar.gz
	color.Blue.Println("\n", string(out))
}
