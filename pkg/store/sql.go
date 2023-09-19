package store

import (
	"database/sql"
	"errors"

	"github.com/andarroyave/reserva-turnos/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func SqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}


	func (s *sqlStore) ReadPatient(ID int) (domain.Patient, error) {
		query := "select * from patients where ID = ?"
		row := s.db.QueryRow(query, ID)
		var patient domain.Patient
		if err := row.Scan(&patient.ID, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.DischargeDate); err != nil {
			return domain.Patient{}, errors.New("patient not found")
		}
		return patient, nil
	}

	func (s *sqlStore) ReadAllPatients() ([]domain.Patient, error) {
		query := "select * from patients"
		sttm, err := s.db.Query(query)
		if err != nil {
			return nil, err
		}
		defer sttm.Close()

		var patients []domain.Patient
		for sttm.Next() {
			var patient domain.Patient
			err := sttm.Scan(&patient.ID, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.DischargeDate)
			if err != nil {
				return nil, err
			}
			patients = append(patients, patient)
		}
		if err = sttm.Err(); err != nil {
			return nil, err
		}
		return patients, nil
	}

	func (s *sqlStore) CreatePatient(patient domain.Patient) error {
		query := "insert into patients (dni, name, lastname, address, dischargedate) values (?, ?, ?, ?, ?)"
		sttm, err := s.db.Prepare(query)
		if err != nil {
			return err
		}
		defer sttm.Close()
		res, err := sttm.Exec(patient.DNI, patient.Name, patient.LastName, patient.Address, patient.DischargeDate)
		if err != nil {
			return err
		}
		_, err = res.RowsAffected()
		if err != nil {
			return err
		}
		return nil
	}

	func (s *sqlStore) UpdatePatient(patient domain.Patient) error {
		query := "update patients set dni = ?, name = ?, LastName = ?, address = ?, dischargedate = ? where ID = ?"
		sttm, err := s.db.Prepare(query)
		if err != nil {
			return err
		}
		defer sttm.Close()

		_, err = sttm.Exec(patient.DNI, patient.Name, patient.LastName, patient.Address, patient.DischargeDate, patient.ID)
		if err != nil {
			return err
		}

		return nil
	}

	func (s *sqlStore) DeletePatient(ID int) error {
		query := "delete from patients where ID = ?"
		_, err := s.db.Exec(query, ID)
		if err != nil {
			return err
		}
		return nil
	}

	func (s *sqlStore) ExistsPatient(dni string) bool {
		query := "select ID_patient from patients where dni like ?"
		row := s.db.QueryRow(query, dni)
		var ID int
		if err := row.Scan(&ID); err != nil {
			return false
		}

		if ID > 0 {
			return true
		}

		return false
	}

