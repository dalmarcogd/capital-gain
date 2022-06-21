package calculator

type Transaction struct {
	Operation Operation `json:"operation"`
	UnitCost  float64   `json:"unit-cost"`
	Quantity  int       `json:"quantity"`
}

type TransactionReturn struct {
	Tax   float64
	Error error
}
