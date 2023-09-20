package domain

type Turn struct {
	Id          int64   `json:"id"`
	Dentist     Dentist `json:"dentist"`
	Patient     Patient `json:"patient"`
	DateHour    string  `json:"datehour"`
	Description string  `json:"description"`
}