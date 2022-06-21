package cliact

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dalmarcogd/capital-gain/internal/calculator"
)

type (
	transaction struct {
		Operation string  `json:"operation"`
		UnitCost  float64 `json:"unit-cost"`
		Quantity  int     `json:"quantity"`
	}

	transactionReturn struct {
		Tax   *float64 `json:"tax,omitempty"`
		Error string   `json:"error,omitempty"`
	}
)

func CalcAction(calc calculator.Calculator) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Bytes()

		if string(data) == "exit" {
			break
		}

		str, err := calcAction(calc, data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(str)
	}
}

func calcAction(calc calculator.Calculator, data []byte) (string, error) {
	var trxs []transaction
	err := json.Unmarshal(data, &trxs)
	if err != nil {
		return "", err
	}

	if len(trxs) > 0 {
		calcTrxs := make([]calculator.Transaction, 0, len(trxs))
		for _, trx := range trxs {
			calcTrxs = append(calcTrxs, calculator.Transaction{
				Operation: calculator.Operation(trx.Operation),
				UnitCost:  trx.UnitCost,
				Quantity:  trx.Quantity,
			})
		}

		transactionReturns := calc.Calc(calcTrxs)
		trxReturn := make([]transactionReturn, 0, len(transactionReturns))

		for _, t := range transactionReturns {
			var tr transactionReturn
			tax := t.Tax

			if t.Error != nil {
				tr.Error = t.Error.Error()
			} else {
				tr.Tax = &tax
			}

			trxReturn = append(trxReturn, tr)
		}

		data, err = json.Marshal(&trxReturn)
		if err != nil {
			return "", err
		}
		
		return string(data), nil
	}

	return "", nil
}
