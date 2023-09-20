package store

import (
	"database/sql"
	"errors"
	"fmt"

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

func (s *sqlStore) ReadPatient(id int) (domain.Patient, error) {
	query := "select * from patients where id = ?"
	row := s.db.QueryRow(query, id)
	var patient domain.Patient
	if err := row.Scan(&patient.Id, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.DischargeDate); err != nil {
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
		err := sttm.Scan(&patient.Id, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.DischargeDate)
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
	query := "update patients set dni = ?, name = ?, lastname = ?, address = ?, dischargedate = ? where id = ?"
	sttm, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer sttm.Close()

	_, err = sttm.Exec(patient.DNI, patient.Name, patient.LastName, patient.Address, patient.DischargeDate, patient.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) DeletePatient(id int) error {
	query := "delete from patients where id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) ExistsPatient(dni string) bool {
	query := "select id_patient from patients where dni like ?"
	row := s.db.QueryRow(query, dni)
	var id int
	if err := row.Scan(&id); err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}

// SQL DENTIST
func (r *sqlStore) GetByIdDentist(id int) (domain.Dentist, error) {
	var dentist domain.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE id = %d;", id)
	row := r.db.QueryRow(query)
	err := row.Scan(&dentist.Id, &dentist.Registration, &dentist.LastName, &dentist.Name)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *sqlStore) GetAllDentist() ([]domain.Dentist, error) {
	var dentists []domain.Dentist

	query := "SELECT * FROM dentists;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(&dentist.Id, &dentist.Registration, &dentist.LastName, &dentist.Name)
		if err != nil {
			return nil, err
		}
		dentists = append(dentists, dentist)
	}

	return dentists, nil
}

func (r *sqlStore) CreateDentist(dentist domain.Dentist) (domain.Dentist, error) {
	query := fmt.Sprintf("INSERT INTO dentists (Registration, lastname, firstname) VALUES ('%s', '%s', '%s');",
		dentist.Registration, dentist.LastName, dentist.Name)
	_, err := r.db.Exec(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *sqlStore) ModifyByIdDentist(id int, dentist domain.Dentist) (domain.Dentist, error) {
	query := fmt.Sprintf("UPDATE dentists SET Registration = '%s', lastname = '%s', firstname = '%s' WHERE id = %d;",
		dentist.Registration, dentist.LastName, dentist.Name, id)
	_, err := r.db.Exec(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *sqlStore) DeleteByIdDentist(id int) error {
	query := fmt.Sprintf("DELETE FROM dentists WHERE id = %d;", id)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
