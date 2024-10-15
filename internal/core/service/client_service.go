package application

import (
	"errors"

	"github.com/pwkm/client/internal/domain"
	"github.com/pwkm/client/internal/ports"
)

var (
	Error_Empty_Mail = errors.New("A email address is needed")
)

type ClientService struct {
	repo ports.DBRepository
}

func NewClientService(repo ports.DBRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateProfile( id string, name string, email string, adress *Address
	Login  *Login) error {
	// Check Gegevens

	// Create Client


	// Save Client gegevens
	return s.repo.Save(client)
}

func (s *ClientService) GetClient(id string) (*domain.Client, error) {
	return s.repo.FindByID(id)
}

func (s *ClientService) RegisterClient(email string, password string) error {
	// Check information

	// Create login enity
	login, err := domain.NewLogin(email, password)
	if err != nil {
		return Error_Empty_Mail
	}

	// Save to the Repository
	return s.repo.SaveLogin(login)
}

// func (s *ClientService) UpdateClient(id string) (*domain.Client, error) {
// 	return s.repo.UpdateC(id)
// }
