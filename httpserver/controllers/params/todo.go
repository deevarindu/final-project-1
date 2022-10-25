package params

type TodoCreateRequest struct {
	Title  string `validate:"required" json:"title"`
	Status string `json:"status"`
}

type TodoUpdateRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
