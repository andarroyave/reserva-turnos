package domain

type Dentist struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Registration string `json:"registration"`
}
