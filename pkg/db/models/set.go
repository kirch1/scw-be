package models

type Set struct {
	ID    int64  `json:"id"`
	Sport string `json:"sport"`
	Brand string `json:"brand"`
	Year  int64  `json:"year"`
	Set   string `json:"set"`
}
