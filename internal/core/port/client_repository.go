package ports

import "github.com/pwkm/client/internal/core/domain"

type DBRepository interface {
	SaveClient(client *domain.Client) error
	FindByID(id string) (*domain.Client, error)
	SaveLogin(login *domain.Login) error
}
