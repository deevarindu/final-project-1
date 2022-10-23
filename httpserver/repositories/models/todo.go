package models

type Todo struct {
	ID     string `example:"1" json:"id"`
	Title  string `example:"do final project" json:"title"`
	Status string `example:"done" json:"status"`
}
