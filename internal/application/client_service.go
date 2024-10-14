package application

import (
	"github.com/pwkm/client/internal/domain"
	"github.com/pwkm/client/internal/ports"
)

type ClientService struct {
	repo ports.ClientRepository
}

func NewClientService(repo ports.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateProfile(client *domain.Client) error {
	return s.repo.Save(client)
}

func (s *ClientService) GetClient(id string) (*domain.Client, error) {
	return s.repo.FindByID(id)
}

func (s *ClientService) RegisterClient(login *domain.Login) error {
	return s.repo.RegisterClient(login)
}

// func (s *ClientService) UpdateClient(id string) (*domain.Client, error) {
// 	return s.repo.UpdateC(id)
// }
