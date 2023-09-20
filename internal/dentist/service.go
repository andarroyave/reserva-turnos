package dentist

import (
	"github.com/andarroyave/reserva-turnos/internal/domain"
)

// Repository es una interfaz que utilizamos para indicar cómo implementar
// un repositorio para Dentistas.
type Repository interface {
	GetByIdDentist(id int) (domain.Dentist, error)
	GetAllDentists() ([]domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error)
	DeleteByIdDentist(id int) error
}

// Service proporciona todas las funcionalidades relacionadas con los Dentistas.
type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetById(id int) (domain.Dentist, error) {
	return s.repository.GetByIdDentist(id)
}

func (s *Service) GetAll() ([]domain.Dentist, error) {
	return s.repository.GetAllDentists()
}

func (s *Service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.CreateDentist(dentist)
}

func (s *Service) ModifyById(id int, dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.ModifyByIdDentist(id, dentist)
}

func (s *Service) DeleteById(id int) error {
	return s.repository.DeleteByIdDentist(id)
}
