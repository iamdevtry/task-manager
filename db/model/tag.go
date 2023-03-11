package model

type Tag struct {
	Id    int64  `json:"id" db:"ID"`
	Title string `json:"title" db:"TITLE"`
	Slug  string `json:"slug" db:"SLUG"`
}
