package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dalmarcogd/capital-gain/internal/calculator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "capital-gain",
		Usage: "calculate stock transaction tax",
		Commands: []*cli.Command{
			{
				Name:  "calculator",
				Usage: "write the transactions inside an array and press enter, type exit to stop using",
				Action: func(ctx *cli.Context) error {
					fmt.Println(ctx.Command.Usage)
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						data := scanner.Bytes()

						if string(data) == "exit" {
							break
						}

						var trxs []calculator.Transaction
						err := json.Unmarshal(data, &trxs)
						if err != nil {
							fmt.Println(err)
							continue
						}

						if len(trxs) > 0 {
							transactionTaxes := calculator.NewCalculator().Calc(trxs)
							data, err = json.Marshal(&transactionTaxes)
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Println(string(data))
						}
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
