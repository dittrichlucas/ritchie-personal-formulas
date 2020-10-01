/*
 * TODOS:
 *	1. Validar as entradas para que sejam no fomato 00h00m
 *	2. Mudar a cor da msg de retorno de acordo com o resultado do dia
 *		Verde: cumpriu o expediente com 10m pra mais ou pra menos
 *		Amarelo: fez horas extras
 *		Vermelho: faltou horas
 *	3. Como saída, desejo que mostre na tela:
 *		- O horário do primeiro período
 *		- O horário do segundo período
 *		- Se fez hora extra, a quantidade de hora extra
 *		- Se as horas trabalhadas forem menor do que o esperado, apresentar quantas horas vão para banco
 *	4. Tornar o retorno mais interativo através do uso de delay e mensagens mais humanas (exemplo do café - Ritchie)
 *	. Aceitar um input de quantidade de horas/dia
 *		Aqui eu pretendo deixar a fórmula flexível, pois cada empresa tem uma quantidade de horas/dia que devem ser cumpridas
 */

package formula

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Formula is...
type Formula struct {
	FirstEntry  string
	FirstExit   string
	SecondEntry string
	SecondExit  string
}

var (
	baseTime, _      = time.ParseDuration("08h48m")
	toleranceTime, _ = time.ParseDuration("0h10m")
)

// Run is...
func (f Formula) Run(writer io.Writer) {
	firstEntry, _ := time.ParseDuration(f.FirstEntry)
	firstExit, _ := time.ParseDuration(f.FirstExit)
	secondEntry, _ := time.ParseDuration(f.SecondEntry)
	secondExit, _ := time.ParseDuration(f.SecondExit)

	firstPeriod := (firstExit - firstEntry).String()
	secondPeriod := (secondExit - secondEntry).String()
	total := ((firstExit - firstEntry) + (secondExit - secondEntry))

	firstPeriod = strings.Replace(firstPeriod, "0s", "", 1)
	secondPeriod = strings.Replace(secondPeriod, "0s", "", 1)
	sTotal := strings.Replace(total.String(), "0s", "", 1)

	color.White(fmt.Sprintf("\n🕗 First period: %s\n", firstPeriod))
	color.White(fmt.Sprintf("🕗 Second period: %s\n\n", secondPeriod))

	if baseTime-toleranceTime <= total && total < baseTime {
		color.Yellow(fmt.Sprintf("🕗 Total work (-): %s\n", sTotal))

		return
	}

	if baseTime+toleranceTime >= total && total > baseTime {
		color.Yellow(fmt.Sprintf("🕗 Total work (+): %s\n", sTotal))

		return
	}

	if baseTime+toleranceTime < total {
		overtimes := (total - baseTime)
		sOvertimes := strings.Replace(overtimes.String(), "0s", "", 1)

		color.Yellow(fmt.Sprintf("🕗 Total work (+): %s\n", sTotal))
		color.Green(fmt.Sprintf("🕗 Overtimes: %s\n\n", sOvertimes))

		return
	}

	if baseTime-toleranceTime > total {
		hourbank := (baseTime - total)
		sHourBank := strings.Replace(hourbank.String(), "0s", "", 1)

		color.Red(fmt.Sprintf("🕗 Total work (-): %s", sTotal))
		color.Red(fmt.Sprintf("🕗 Hour bank: %s\n\n", sHourBank))

		return
	}

	color.Green(fmt.Sprintf("🕗 Total work: %s\n", sTotal))
}
