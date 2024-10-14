package ports

import "github.com/pwkm/client/internal/domain"

type ClientRepository interface {
	Save(client *domain.Client) error
	FindByID(id string) (*domain.Client, error)
	RegisterClient(login *domain.Login) error
}
