// // This is the formula implementation class.
// // Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// )

// // Formula is
// type Formula struct{}

// /*
// 	TODO:
// 		1. Criar um arquivo separado para gerenciar o output
// 		2. Criar métodos de: criar, ler, listar, escrever e remover
// 		3. Alterar a fórmula para chamar os métodos criados no passo anterior
// */

// // Run is
// func (f Formula) Run(writer io.Writer) {
// 	fTest := "rit math sum --number_1=35 --number_2=50"

// 	runFormula(fTest)

// 	// cmd := exec.Command("rit", "math", "sum", "--number_1=45", "--number_2=10")
// 	// cmd.Run()

// 	pwd := RitchieDir()

// 	file, _ := ioutil.ReadFile(pwd + OutputFile)

// 	var data output

// 	if err := json.Unmarshal(file, &data); err != nil {
// 		panic(err)
// 	}

// 	n1, _ := strconv.Atoi(data.Output)
// 	multiply := n1 * 5

// 	fmt.Println(multiply)
// 	// defer RemoveFile()

// 	in := make(map[string]interface{})
// 	out := make(map[string]interface{})

// 	// in["input_1"] = "245"
// 	out["output_1"] = "kaduzera eh fera"

// 	if err := WriteFile(in, out); err != nil {
// 		os.Exit(1)
// 	}
// }

// // RemoveFile is
// func RemoveFile() {
// 	dir := RitchieDir()
// 	file := filepath.Join(dir, OutputFile)
// 	os.Remove(file)
// }
