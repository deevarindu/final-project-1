package views

type Response struct {
	ID     int    `example:"1" json:"id"`
	Title  string `example:"Mengerjakan final project" json:"title"`
	Status string `example:"in progress" json:"status"`
}
