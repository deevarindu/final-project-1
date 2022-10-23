package models

type Todo struct {
	ID     int    `example:"1" json:"id"`
	Title  string `example:"do final project" json:"title"`
	Status string `example:"done" json:"status"`
}
