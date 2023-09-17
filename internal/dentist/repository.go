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

func (s *SqlStore) GetByMatriculation(matriculation int) (domain.Dentist, error) {
	var dentistReturn domain.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE matriculation = %d;", matriculation)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.Matriculation, &dentistReturn.LastName, &dentistReturn.Name)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) Modify(matriculation int, dentist domain.Dentist) (domain.Dentist, error) {
	query := fmt.Sprintf("UPDATE dentists SET lastname = '%s', firstname = '%s' WHERE matriculation = %d;",
		dentist.LastName, dentist.Name, matriculation)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}
