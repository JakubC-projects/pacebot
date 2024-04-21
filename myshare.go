package peacefulroad

import (
	"context"
)

const BukUstronClub = 3982

type Status struct {
	TransactionsAmount float64 `json:"transactionsAmount"`
	Target             float64 `json:"target"`
	Currency           string  `json:"currency"`
}

type Org struct {
}

type StatusService interface {
	GetStatus(ctx context.Context, user User) (Status, error)
}

type OrgService interface {
	GetOrgs(ctx context.Context, user User) ([]Org, error)
}
