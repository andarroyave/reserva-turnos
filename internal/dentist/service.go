package dentist

import (
	"github.com/andarroyave/reserva-turnos/internal/domain"
)

// Repository es una interfaz que utilizamos para indicar cómo implementar
// un repositorio para Dentistas.
type Repository interface {
	GetByMatriculation(matriculation int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Modify(matriculation int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(matriculation int) error
}

// Service proporciona todas las funcionalidades relacionadas con los Dentistas.
type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByMatriculation(matriculation int) (domain.Dentist, error) {
	return s.repository.GetByMatriculation(matriculation)
}

func (s *Service) GetAll() ([]domain.Dentist, error) {
	return s.repository.GetAll()
}

func (s *Service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.Create(dentist)
}

func (s *Service) ModifyByMatriculation(matriculation int, dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.Modify(matriculation, dentist)
}

func (s *Service) DeleteByMatriculation(matriculation int) error {
	return s.repository.Delete(matriculation)
}
