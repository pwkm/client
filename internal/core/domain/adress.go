package domain

import (
	"github.com/google/uuid"
)

// var (
// 	err = errors.New("currency mismatch")
// )

type Address struct {
	ID        uuid.UUID
	Street    string
	Number    string
	Postcode  int
	Community string
}

func NewAddress(id uuid.UUID, street string, number string, postcode int, community string) (*Address, error) {
	return &Address{
		ID:        id,
		Street:    street,
		Number:    number,
		Postcode:  postcode,
		Community: community,
	}, nil
}
