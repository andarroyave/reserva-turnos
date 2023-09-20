package store

import "github.com/andarroyave/reserva-turnos/internal/domain"

type StoreInterface interface {
	ReadPatient(id int) (domain.Patient, error)
	ReadAllPatients() ([]domain.Patient, error)
	CreatePatient(patient domain.Patient) error
	UpdatePatient(patient domain.Patient) error
	DeletePatient(id int) error
	ExistsPatient(dni string) bool
	//Dentist
	GetByIdDentist(id int) (domain.Dentist, error)
	GetAllDentist() ([]domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error)
	DeleteByIdDentist(id int) error
}
