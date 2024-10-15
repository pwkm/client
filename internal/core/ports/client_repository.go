package ports

import "github.com/pwkm/client/internal/domain"

type DBRepository interface {
	Save(client *domain.Client) error
	FindByID(id string) (*domain.Client, error)
	SaveLogin(login *domain.Login) error
}
