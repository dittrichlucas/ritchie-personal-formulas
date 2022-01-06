package formula

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// )

// type OutputJSON map[string]interface{}

// var outputFile = getOutputPath()

// func RunFormula(command string) error {
// 	cmdSplit := strings.Split(command, " ")
// 	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
// 	if err := cmd.Start(); err != nil {
// 		return err
// 	}

// 	cmd.Wait()

// 	exitCode := cmd.ProcessState.ExitCode()
// 	fmt.Println(exitCode)
// 	if exitCode > 0 {
// 		return errors.New("command not executed")
// 	}

// 	return nil
// }

// // Validar os comportamentos quando o arquivo existe e quando não existe
// func SetOutput(key string, value interface{}) error {
// 	file, _ := ioutil.ReadFile(outputFile)

// 	var data OutputJSON
// 	if err := json.Unmarshal(file, &data); err != nil {
// 		return err
// 	}

// 	data[key] = value

// 	outputJSON, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile(outputFile, outputJSON, os.ModePerm); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Validar se o arquivo existe

// func GetOutput(key string) (interface{}, error) {
// 	file, _ := ioutil.ReadFile(outputFile)

// 	var data OutputJSON
// 	if err := json.Unmarshal(file, &data); err != nil {
// 		return nil, err
// 	}

// 	return data[key], nil
// }

// func RemoveOutput(key string) error {
// 	file, _ := ioutil.ReadFile(outputFile)

// 	var data OutputJSON
// 	if err := json.Unmarshal(file, &data); err != nil {
// 		return err
// 	}
// 	delete(data, key)

// 	outputJSON, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile(outputFile, outputJSON, os.ModePerm); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func getOutputPath() string {
// 	myDir, _ := os.UserHomeDir()

// 	return filepath.Join(myDir, ".rit", "output.json")
// }

// func parseData(isRemove bool) {
// 	file, _ := ioutil.ReadFile(outputFile)

// 	var data OutputJSON
// 	if err := json.Unmarshal(file, &data); err != nil {
// 		return err
// 	}
// 	delete(data, key)

// 	outputJSON, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile(outputFile, outputJSON, os.ModePerm); err != nil {
// 		return err
// 	}

// 	return nil
// }
