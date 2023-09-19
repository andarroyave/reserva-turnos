package turn

import (
	"github.com/andarroyave/reserva-turnos/internal/domain"
)

type IService interface {
	GetTurnById(id int64) (*domain.Turn, error)
	CreateTurn(t domain.Turn) (*domain.Turn, error)
	UpdateTurn(t domain.Turn) (*domain.Turn, error)
	UpdateTurnFields(t domain.Turn) (*domain.Turn, error)
	DeleteTurn(id int64) (string , error)
	GetTurnByDNI(dni string) ([]domain.Turn, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetTurnById(id int64) (*domain.Turn, error) {
	turn, err := s.Repository.GetTurnById(id)
	if err != nil {
		return nil, err
	}
	return turn, nil
}

func (s *Service) CreateTurn(t domain.Turn) (*domain.Turn, error) {
	turn, err := s.Repository.CreateTurn(t)
	if err != nil {
		return nil, err
	}
	return turn, nil
}

func (s *Service) UpdateTurn(t domain.Turn) (*domain.Turn, error) {
	turn, err := s.Repository.UpdateTurn(t)
	if err != nil {
		return nil, err
	}
	return turn, nil
}

func (s *Service) UpdateTurnFields(t domain.Turn) (*domain.Turn, error) {
	turn, err := s.Repository.UpdateTurnFields(t)
	if err != nil {
		return nil, err
	}
	return turn, nil
}

func (s *Service) DeleteTurn(id int64) (string , error) {
	res, err := s.Repository.DeleteTurn(id)
	if err != nil {
		return "", err
	}
	return res, nil
}


func (s *Service) GetTurnByDNI(dni string) ([]domain.Turn, error) {
	turns, err := s.Repository.GetTurnByDNI(dni)
	if err != nil {
		return nil, err
	}
	return turns, nil
}