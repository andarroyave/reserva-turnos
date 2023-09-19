package store

import "github.com/andarroyave/reserva-turnos/internal/domain"

type StoreInterface interface {
	ReadPatient(id int) (domain.Patient, error)
	ReadAllPatients() ([]domain.Patient, error)
	CreatePatient(patient domain.Patient) error
	UpdatePatient(patient domain.Patient) error
	DeletePatient(ID int) error
	ExistsPatient(dni string) bool
}