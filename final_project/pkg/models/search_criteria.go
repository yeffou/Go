package models

type SearchCriteria struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   int    `json:"year,omitempty"`
}
