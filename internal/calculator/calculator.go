package calculator

import "math"

const (
	TaxPercentage  = 0.2
	OperationLimit = 20000.00
)

type Calculator interface {
	Calc(trxs []Transaction) []TransactionTax
}

type calculator struct{}

func NewCalculator() Calculator {
	return calculator{}
}

func (c calculator) Calc(trxs []Transaction) []TransactionTax {
	trxsTax := make([]TransactionTax, len(trxs))
	var totalQuantity int
	var priceAverage float64
	var accumulatedLoss float64
	for i, trx := range trxs {
		var tax float64

		switch trx.Operation {
		case Buy:
			// calculate the price average
			priceAverage = math.Round(((float64(totalQuantity)*priceAverage)+(float64(trx.Quantity)*trx.UnitCost))/(float64(totalQuantity+trx.Quantity))*100) / 100
			totalQuantity += trx.Quantity
		case Sell:
			totalQuantity -= trx.Quantity
			// calculate profit vs loss
			if trx.UnitCost*float64(trx.Quantity) > OperationLimit && trx.UnitCost > priceAverage {
				profit := (trx.UnitCost * float64(trx.Quantity)) - (priceAverage * float64(trx.Quantity))

				// reduce profit loss
				if accumulatedLoss > 0 {
					if profit > accumulatedLoss {
						profit -= accumulatedLoss
					} else {
						accumulatedLoss -= profit
						profit = 0
					}
				}

				if profit > 0 {
					tax = math.Round((profit*TaxPercentage)*100) / 100
				}
			} else {
				// accumulate the loss of the operation
				accumulatedLoss += (priceAverage * float64(trx.Quantity)) - (trx.UnitCost * float64(trx.Quantity))
			}
		}

		trxsTax[i] = TransactionTax{Tax: tax}
	}

	return trxsTax
}
