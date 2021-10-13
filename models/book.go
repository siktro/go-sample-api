package models

type Book struct {
	ID     int     `json:"id,omitempty"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   *uint16 `json:"year"`
}
