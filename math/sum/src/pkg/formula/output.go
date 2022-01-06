package formula

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type OutputJSON map[string]interface{}

var outputFile = getOutputPath()

// TODO: precisa documentar bem as funções para auxiliar os usuários

func RunFormula(command string) error {
	cmdSplit := strings.Split(command, " ")
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	if err := cmd.Start(); err != nil {
		return err
	}

	cmd.Wait()

	exitCode := cmd.ProcessState.ExitCode()
	if exitCode > 0 {
		return errors.New("command not executed")
	}

	return nil
}

// Validar os comportamentos quando o arquivo existe e quando não existe
func SetOutput(key string, value interface{}) error {
	isFileExist := fileExists(outputFile)
	if isFileExist {
		data, err := parseData()
		if err != nil {
			fmt.Println(err)
			return err
		}

		data[key] = value
		if err := writeData(data); err != nil {
			return err
		}
	} else {
		data := make(OutputJSON)
		data[key] = value
		if err := writeData(data); err != nil {
			return err
		}
	}

	return nil
}

// Validar se o arquivo existe
func GetOutput(key string) (interface{}, error) {
	data, err := parseData()
	if err != nil {
		return nil, err
	}

	return data[key], nil
}

func RemoveOutput(key string) error {
	data, err := parseData()
	if err != nil {
		return err
	}

	delete(data, key)

	if err := writeData(data); err != nil {
		return err
	}

	return nil
}

func getOutputPath() string {
	myDir, _ := os.UserHomeDir()

	return filepath.Join(myDir, ".rit", "output.json")
}

func parseData() (OutputJSON, error) {
	var data OutputJSON

	file, err := ioutil.ReadFile(outputFile)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return data, err
	}

	return data, nil
}

func writeData(data OutputJSON) error {
	outputJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outputFile, outputJSON, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
