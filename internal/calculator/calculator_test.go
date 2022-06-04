package calculator

import (
	"context"
	"reflect"
	"testing"
)

func Test_calculator_Calc(t *testing.T) {
	type args struct {
		ctx  context.Context
		trxs []Transaction
	}
	var tests = []struct {
		name string
		args args
		want []TransactionTax
	}{
		{
			name: "case 1",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  1000,
					},
					{
						Operation: Sell,
						UnitCost:  15.00,
						Quantity:  50,
					},
					{
						Operation: Sell,
						UnitCost:  15.00,
						Quantity:  50,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			name: "case 2",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  5.00,
						Quantity:  5000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 10000.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			name: "case 3",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  5.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  3000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 1000.00,
				},
			},
		},
		{
			name: "case 4",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Buy,
						UnitCost:  25.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  15.00,
						Quantity:  10000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			name: "case 5",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Buy,
						UnitCost:  25.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  15.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  25.00,
						Quantity:  5000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 10000.00,
				},
			},
		},
		{
			name: "case 6",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  2.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  2000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  2000,
					},
					{
						Operation: Sell,
						UnitCost:  25.00,
						Quantity:  1000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 3000.00,
				},
			},
		},
		{
			name: "case 7",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  2.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  2000,
					},
					{
						Operation: Sell,
						UnitCost:  20.00,
						Quantity:  2000,
					},
					{
						Operation: Sell,
						UnitCost:  25.00,
						Quantity:  1000,
					},
					{
						Operation: Buy,
						UnitCost:  20.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  15.00,
						Quantity:  5000,
					},
					{
						Operation: Sell,
						UnitCost:  30.00,
						Quantity:  4350,
					},
					{
						Operation: Sell,
						UnitCost:  30.00,
						Quantity:  650,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 3000.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 3700.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			name: "case 8",
			args: args{
				ctx: context.Background(),
				trxs: []Transaction{
					{
						Operation: Buy,
						UnitCost:  10.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  50.00,
						Quantity:  10000,
					},
					{
						Operation: Buy,
						UnitCost:  20.00,
						Quantity:  10000,
					},
					{
						Operation: Sell,
						UnitCost:  50.00,
						Quantity:  10000,
					},
				},
			},
			want: []TransactionTax{
				{
					Tax: 0.00,
				},
				{
					Tax: 80000.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 60000.00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calculator{}
			got := c.Calc(tt.args.trxs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
