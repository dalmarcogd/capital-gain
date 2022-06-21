package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dalmarcogd/capital-gain/internal/calculator"
	"github.com/dalmarcogd/capital-gain/internal/cliact"
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
					cliact.CalcAction(calculator.NewCalculator())
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
