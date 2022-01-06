/*
 * TODOS:
 *	1. Validar se o valor passado contem apenas números
 */

package formula

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

// Input is my stdin
type Input struct {
	Zip string
}

// Address is body response struct
type Address struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
	Unidade     string
	Ibge        string
	Gia         string
	Erro        bool
}

const zipErrorLength = "Zip code must be 8 characters. Please try again."
const invalidZipCodeError = "Invalid zip code! Please try again."

// Run is my runner
func (in Input) Run() {
	zip := strings.Replace(in.Zip, "-", "", -1)

	if len(zip) != 8 {
		color.Red(fmt.Sprintln(zipErrorLength))
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zip)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err, zip)
	}

	defer resp.Body.Close()

	var address Address

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal([]byte(body), &address)

		if address.Erro {
			color.Red(fmt.Sprintln(invalidZipCodeError))
			return
		}

		color.Green(fmt.Sprintf("Zip code: %v\nPublic place: %v\nNeighborhood: %v\nLocality: %v\nUF: %v\n",
			address.Cep, address.Logradouro, address.Bairro, address.Localidade, address.Uf))
	}
}
