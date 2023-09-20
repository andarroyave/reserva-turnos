package patient

import "github.com/andarroyave/reserva-turnos/internal/domain"

type Service interface {
	GetById(Id int) (domain.Patient, error)
	GetAllPatients() ([]domain.Patient, error)
	CreatePatient(p domain.Patient) (domain.Patient, error)
	DeletePatient(Id int) error
	UpdatePatient(Id int, u domain.Patient) (domain.Patient, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetById(Id int) (domain.Patient, error) {
	p, err := s.r.GetById(Id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s service) GetAllPatients() ([]domain.Patient, error) {
	p, err := s.r.GetAllPatients()
	if err != nil {
		return []domain.Patient{}, err
	}
	return p, nil
}

func (s service) CreatePatient(p domain.Patient) (domain.Patient, error) {
	//fmt.Println(p)
	p, err := s.r.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) DeletePatient(Id int) error {
	err := s.r.DeletePatient(Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdatePatient(Id int, u domain.Patient) (domain.Patient, error) {
	p, err := s.r.GetById(Id)
	if err != nil {
		return domain.Patient{}, err
	}
	if u.DNI != "" {
		p.DNI = u.DNI
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.LastName != "" {
		p.LastName = u.LastName
	}
	if u.Address != "" {
		p.Address = u.Address
	}
	if u.DischargeDate != "" {
		p.DischargeDate = u.DischargeDate
	}
	p, err = s.r.UpdatePatient(p.DNI, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}