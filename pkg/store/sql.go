package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/andarroyave/reserva-turnos/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) ReadPatient(Id int) (domain.Patient, error) {
	query := "select * from patients where Id = ?"
	row := s.DB.QueryRow(query, Id)
	var patient domain.Patient
	if err := row.Scan(&patient.Id, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.DischargeDate); err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}

func (s *SqlStore) ReadAllPatients() ([]domain.Patient, error) {
	query := "select * from patients"
	sttm, err := s.DB.Query(query)
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

func (s *SqlStore) CreatePatient(patient domain.Patient) error {
	query := "insert into patients (dni, name, lastname, address, dischargedate) values (?, ?, ?, ?, ?)"
	sttm, err := s.DB.Prepare(query)
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

func (s *SqlStore) UpdatePatient(patient domain.Patient) error {
	query := "update patients set dni = ?, name = ?, LastName = ?, address = ?, dischargedate = ? where Id = ?"
	sttm, err := s.DB.Prepare(query)
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

func (s *SqlStore) DeletePatient(Id int) error {
	query := "delete from patients where Id = ?"
	_, err := s.DB.Exec(query, Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) ExistsPatient(dni string) bool {
	query := "select Id_patient from patients where dni like ?"
	row := s.DB.QueryRow(query, dni)
	var Id int
	if err := row.Scan(&Id); err != nil {
		return false
	}

	if Id > 0 {
		return true
	}

	return false
}

func (s *SqlStore) GetTurnById(Id int64) (*domain.Turn, error) {
	var turn domain.Turn
	var patientId int64
	var dentistId int64
	queryT := "SELECT * FROM Turns WHERE id = ?;"
	rowT := s.DB.QueryRow(queryT, Id)
	errT := rowT.Scan(&turn.Id, &patientId, &dentistId, &turn.DateHour, &turn.Description)
	if errT != nil {
		return nil, errT
	}
	queryP := "SELECT * FROM Patients WHERE id = ?;"
	rowP := s.DB.QueryRow(queryP, patientId)
	errP := rowP.Scan(&turn.Patient.Id, &turn.Patient.Name, &turn.Patient.LastName, &turn.Patient.Address, &turn.Patient.DNI, &turn.Patient.DischargeDate)
	if errP != nil {
		return nil, errP
	}
	queryD := "SELECT * FROM Dentists WHERE id = ?;"
	rowD := s.DB.QueryRow(queryD, dentistId)
	errD := rowD.Scan(&turn.Dentist.Id, &turn.Dentist.Name, &turn.Dentist.LastName, &turn.Dentist.Registration)
	if errD != nil {
		return nil, errD
	}

	return &turn, nil
}

func (s *SqlStore) CreateTurn(turn domain.Turn) (*domain.Turn, error) {

	query := "INSERT INTO Turns (PatientId, DentistId, DateHour, Description)	VALUES (?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(turn.Patient.Id, turn.Dentist.Id, turn.DateHour, turn.Description)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	turn.Id = int64(id)
	return &turn, nil

}

func (s *SqlStore) UpdateTurn(t domain.Turn) (*domain.Turn, error) {
	var turnId int64 = t.Id
	var patientId int64
	var dentistId int64
	var turn domain.Turn
	query := "SELECT * FROM Turns WHERE id = ?;"
	row := s.DB.QueryRow(query, turnId)
	err := row.Scan(&turn.Id, &patientId, &dentistId, &turn.DateHour, &turn.Description)
	if err != nil {
		return nil, err
	}

	queryI := "UPDATE Turns SET PatientId = ?, DentistId = ?, DateHour = ?, Description = ? WHERE id = ?;"
	stmt, err := s.DB.Prepare(queryI)
	if err != nil {
		return nil, err
	}
	println(t.Patient.Id, t.Dentist.Id, t.DateHour, t.Description, t.Id)
	res, err := stmt.Exec(t.Patient.Id, t.Dentist.Id, t.DateHour, t.Description, t.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Print(res)
	return &t, nil
}

func (s *SqlStore) UpdateTurnFields(t domain.Turn) (*domain.Turn, error) {
	var turnId int64 = t.Id
	var patientId int64
	var dentistId int64
	var turn domain.Turn
	query := "SELECT * FROM Turns WHERE id = ?;"
	row := s.DB.QueryRow(query, turnId)
	err := row.Scan(&turn.Id, &patientId, &dentistId, &turn.DateHour, &turn.Description)
	if err != nil {
		return nil, err
	}
	turn.Patient.Id = patientId
	turn.Dentist.Id = dentistId
	fmt.Print("hola")
	if t.Patient.Id != 0 {
		turn.Patient = t.Patient
	}
	if t.Dentist.Id != 0 {
		turn.Dentist = t.Dentist
	}
	if t.DateHour != "" {
		turn.DateHour = t.DateHour
	}
	if t.Description != "" {
		turn.Description = t.Description
	}
	queryI := "UPDATE Turns SET PatientId = ?, DentistId = ?, DateHour = ?, Description = ? WHERE id = ?;"
	stmt, err := s.DB.Prepare(queryI)
	if err != nil {
		return nil, err
	}
	println(turn.Patient.Id, turn.Dentist.Id, turn.DateHour, turn.Description, turn.Id)
	res, err := stmt.Exec(turn.Patient.Id, turn.Dentist.Id, turn.DateHour, turn.Description, turn.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Print(res)
	turnRes, err := s.GetTurnById(turnId)
	if err != nil {
		return nil, err
	}
	return turnRes, nil
}

func (s *SqlStore) DeleteTurn(id int64) (string, error) {
	queryD := "DELETE FROM Turns WHERE id = ?;"
	stmt, err := s.DB.Prepare(queryD)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Print(res)
	return "Borrado exitoso", nil
}

func (s *SqlStore) GetTurnByDNI(dni string) ([]domain.Turn, error) {
	var turns []domain.Turn
	query := "SELECT t.Id as t_id, t.PatientId as t_patientid, t.DentistId as t_dentistid, t.DateHour, t.Description, p.Name as p_name, p.LastName as p_lastname, p.Address as p_address, p.DNI as p_dni, p.DischargeDate as p_dischargedate, d.Name as d_name, d.LastName as d_lastname, d.Registration as d_registration FROM Turns t inner join Patients p on t.PatientId = p.Id inner join Dentists d on t.DentistId = d.Id  WHERE p.DNI  = ?;"
	rows, err := s.DB.Query(query, dni)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t domain.Turn
		if err := rows.Scan(&t.Id, &t.Patient.Id, &t.Dentist.Id, &t.DateHour, &t.Description, &t.Patient.Name, &t.Patient.LastName, &t.Patient.Address, &t.Patient.DNI, &t.Patient.DischargeDate, &t.Dentist.Name, &t.Dentist.LastName, &t.Dentist.Registration); err != nil {
			panic(err)
		}
		turns = append(turns, t)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return turns, nil
}
