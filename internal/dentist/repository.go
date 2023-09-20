package dentist

import (
	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/andarroyave/reserva-turnos/pkg/store"
)

type Repository interface {
	GetByIdDentist(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error)
	DeleteByIdDentist(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetByIdDentist(id int) (domain.Dentist, error) {
	dentist, err := r.storage.GetByIdDentist(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) CreateDentist(dentist domain.Dentist) (domain.Dentist, error) {
	createdDentist, err := r.storage.CreateDentist(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return createdDentist, nil
}

func (r *repository) ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error) {
	modifiedDentist, err := r.storage.ModifyByIdDentist(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return modifiedDentist, nil
}

func (r *repository) DeleteByIdDentist(id int) error {
	err := r.storage.DeleteByIdDentist(id)
	if err != nil {
		return err
	}
	return nil
}
