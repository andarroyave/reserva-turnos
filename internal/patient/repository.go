package patient

import (
	"errors"
	"fmt"

	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/andarroyave/reserva-turnos/pkg/store"
)

type Repository interface {
	GetByID(ID int) (domain.Patient, error)
	GetAllPatients() ([]domain.Patient, error)
	CreatePatient(p domain.Patient) (domain.Patient, error)
	DeletePatient(ID int) error
	UpdatePatient(DNI string, p domain.Patient) (domain.Patient, error)
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(ID int) (domain.Patient, error) {
	patient, err := r.storage.ReadPatient(ID)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *repository) GetAllPatients() ([]domain.Patient, error) {
	patients, err := r.storage.ReadAllPatients()
	if err != nil {
		return []domain.Patient{}, errors.New("patients not found")
	}
	return patients, nil
}

func (r *repository) CreatePatient(p domain.Patient) (domain.Patient, error) {
	fmt.Println(p)
	if r.storage.ExistsPatient(p.DNI) {
		return domain.Patient{}, errors.New("patient already exists")
	}
	err := r.storage.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return p, nil
}

func (r *repository) DeletePatient(ID int) error {
	err := r.storage.DeletePatient(ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdatePatient(DNI string, p domain.Patient) (domain.Patient, error) {
	if r.storage.ExistsPatient(p.DNI) {
		return domain.Patient{}, errors.New("dni already exists")
	}
	err := r.storage.UpdatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}