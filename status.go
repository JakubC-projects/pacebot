package pacebot

type Status struct {
	TransactionsAmount float64 `json:"transactionsAmount"`
	PercentageValue    float64 `json:"percentageValue"`
	Target             float64 `json:"target"`
	Currency           string  `json:"currency"`
}
