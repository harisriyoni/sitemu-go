package web

type TypeGaleriCreateRequest struct {
	Type string `validate:"required"`
}

type TypeGaleriUpdateRequest struct {
	Type string `validate:"required"`
}

type TypeGaleriResponse struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	UserID int    `json:"user_id"`
}
