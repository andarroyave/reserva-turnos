package domain

type Dentist struct {
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Matriculation string `json:"matriculation"`
}
type errorMessage struct {
	ErrorInfo string `json:"error"`
}
