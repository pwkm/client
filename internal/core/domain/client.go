package domain

import (
	"github.com/google/uuid"
)

type Client struct {
	ID     uuid.UUID
	Name   string
	Email  string
	Adress *Address
}

func NewClient(id uuid.UUID, name string, email string, address *Address) (*Client, error) {
	return &Client{
		ID:     id,
		Name:   name,
		Email:  email,
		Adress: address,
	}, nil
}
