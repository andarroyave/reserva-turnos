package domain

type Dentist struct {
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Registration string `json:"registration"`
}
type errorMessage struct {
	ErrorInfo string `json:"error"`
}
