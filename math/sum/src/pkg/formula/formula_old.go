package formula

// import (
// 	"encoding/json"
// 	"io"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// )

// // Formula is
// type Formula struct {
// 	Number1 string
// 	Number2 string
// }

// // Output is
// type Output struct {
// 	Input *Formula `json:"input,omitempty"`
// 	Sum   string   `json:"output"`
// }

// const outputFile = "/output.json"

// // Run is
// func (h Formula) Run(writer io.Writer) {
// 	n1, _ := strconv.Atoi(h.Number1)
// 	n2, _ := strconv.Atoi(h.Number2)

// 	sum := n1 + n2
// 	output := strconv.Itoa(sum)
// 	myObj := Output{Sum: output}
// 	bytes, _ := json.Marshal(myObj)

// 	pwd := RitchieDir()

// 	if err := ioutil.WriteFile(pwd+outputFile, bytes, os.ModePerm); err != nil {
// 		os.Exit(1)
// 	}

// }

// // RitchieDir is
// func RitchieDir() string {
// 	myDir, _ := os.UserHomeDir()
// 	return filepath.Join(myDir, ".rit")
// }
