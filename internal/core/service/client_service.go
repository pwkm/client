package application

import (
	"github.com/pwkm/client/internal/core/domain"
	ports "github.com/pwkm/client/internal/core/port"
)

// var (
// 	Error_Empty_Mail = errors.New("A email address is needed")
// )

type ClientService struct {
	repo ports.DBRepository
}

func NewClientService(repo ports.DBRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) RegisterClient(login *domain.Login) error {

	// Save the login to the Repository
	return s.repo.SaveLogin(login)
}

func (s *ClientService) CreateProfile(client *domain.Client) error {

	// Save Client gegevens
	return s.repo.SaveClient(client)
}

func (s *ClientService) GetClient(id string) (*domain.Client, error) {
	return s.repo.FindByID(id)
}

// func (s *ClientService) UpdateClient(id string) (*domain.Client, error) {
// 	return s.repo.UpdateC(id)
// }
