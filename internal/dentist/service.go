package dentist

import (
	"github.com/andarroyave/reserva-turnos/internal/domain"
)

type Service interface {
	GetByIdDentist(Id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	ModifyByIdDentist(Id int, dentist domain.Dentist) (domain.Dentist, error)
	DeleteByIdDentist(Id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetByIdDentist(id int) (domain.Dentist, error) {
	return s.repository.GetByIdDentist(id)
}

func (s *service) CreateDentist(dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.CreateDentist(dentist)
}

func (s *service) ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error) {
	return s.repository.ModifyByIdDentist(id, dentist)
}

func (s *service) DeleteByIdDentist(id int) error {
	return s.repository.DeleteByIdDentist(id)
}
