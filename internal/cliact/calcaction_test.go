package cliact

import (
	"testing"

	"github.com/dalmarcogd/capital-gain/internal/calculator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalcAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculator := calculator.NewMockCalculator(ctrl)

	t.Run("fail deserialize", func(t *testing.T) {
		data := "[{\"operation\":\"buy\", \"unit-cost\": 10, \"quantity\": 10000}, {\"operation\":\"sell\", \"unit-cost\""

		ret, err := calcAction(mockCalculator, []byte(data))
		assert.EqualError(t, err, "unexpected end of JSON input")
		assert.Empty(t, ret)
	})

	t.Run("fail deserialize", func(t *testing.T) {
		data := "[{\"operation\":\"buy\", \"unit-cost\": 10, \"quantity\": 10000}, {\"operation\":\"sell\", \"unit-cost\":20, \"quantity\": 11000}, {\"operation\":\"sell\", \"unit-cost\": 20, \"quantity\": 5000}]"
		dataRet := "[{\"tax\":0},{\"error\":\"can't sell more stocks than you have\"},{\"tax\":10000}]"

		mockCalculator.EXPECT().Calc([]calculator.Transaction{
			{
				Operation: calculator.Buy,
				UnitCost:  10.00,
				Quantity:  10000,
			},
			{
				Operation: calculator.Sell,
				UnitCost:  20.00,
				Quantity:  11000,
			},
			{
				Operation: calculator.Sell,
				UnitCost:  20.00,
				Quantity:  5000,
			},
		}).Return([]calculator.TransactionReturn{
			{
				Tax: 0,
			},
			{
				Error: calculator.ErrInsufficientQuantity,
			},
			{
				Tax: 10000,
			},
		})

		ret, err := calcAction(mockCalculator, []byte(data))
		assert.NoError(t, err)
		assert.Equal(t, dataRet, ret)
	})
}
