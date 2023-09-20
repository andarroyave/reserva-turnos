package domain

type Patient struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Address       string `json:"address"`
	DNI           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}