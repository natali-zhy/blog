package models

type Article struct {
	Id     string `json:"id" db:"id"`
	Author string `json:"Author" db:"author"`
	Text   string `json:"Text" db:"text"`
	Data   string `json:"Data" db:"data"`
}
