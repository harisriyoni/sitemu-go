package web

type GaleriCreateRequest struct {
	TypeGaleriID int    `validate:"required"`
	TitleImage   string `validate:"required"`
}

type GaleriUpdateRequest struct {
	TypeGaleriID int    `validate:"required"`
	TitleImage   string `validate:"required"`
}

type GaleriResponse struct {
	ID           int    `json:"id"`
	TypeGaleriID int    `json:"type_galeri_id"`
	TitleImage   string `json:"title_image"`
	Image        string `json:"image"`
}
