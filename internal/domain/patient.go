package domain

type Patient struct {
	ID            int64  `json:"ID"`
	Name          string `json:"name"`
	LastName      string `json:"LastName"`
	Address       string `json:"address"`
	DNI           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}