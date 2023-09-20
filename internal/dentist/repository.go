package dentist

import (
	"database/sql"
	"fmt"

	"github.com/andarroyave/reserva-turnos/internal/domain"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetById(id int) (domain.Dentist, error) {
	var dentistReturn domain.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.Id, &dentistReturn.Registration, &dentistReturn.LastName, &dentistReturn.Name)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) Create(dentist domain.Dentist) (domain.Dentist, error) {
	query := fmt.Sprintf("INSERT INTO dentists (Registration, lastname, firstname) VALUES ('%s', '%s', '%s');",
		dentist.Registration, dentist.LastName, dentist.Name)
	_, err := s.DB.Exec(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlStore) Modify(id int, dentist domain.Dentist) (domain.Dentist, error) {
	query := fmt.Sprintf("UPDATE dentists SET Registration = '%s', lastname = '%s', firstname = '%s' WHERE id = %d;",
		dentist.Registration, dentist.LastName, dentist.Name, id)
	_, err := s.DB.Exec(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlStore) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM dentists WHERE id = %d;", id)
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
