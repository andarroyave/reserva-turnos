package store

import "github.com/andarroyave/reserva-turnos/internal/domain"

type StoreInterface interface {
	ReadPatient(id int) (domain.Patient, error)
	ReadAllPatients() ([]domain.Patient, error)
	CreatePatient(patient domain.Patient) error
	UpdatePatient(patient domain.Patient) error
	DeletePatient(Id int) error
	ExistsPatient(dni string) bool

	GetTurnById(id int64) (*domain.Turn, error)
	CreateTurn(turn domain.Turn) (*domain.Turn, error)
	UpdateTurn(turn domain.Turn) (*domain.Turn, error)
	UpdateTurnFields(turn domain.Turn) (*domain.Turn, error)
	DeleteTurn(id int64) (string , error)
	GetTurnByDNI(dni string) ([]domain.Turn, error)
}
