package web

type OrganisasiCreateRequest struct {
	UserID  int    `json:"user_id"`
	Jabatan string `json:"jabatan" validate:"required"`
	Nama    string `json:"nama" validate:"required"`
	Image   string `json:"image"`
}

type OrganisasiUpdateRequest struct {
	Jabatan string `json:"jabatan" validate:"required"`
	Nama    string `json:"nama" validate:"required"`
	Image   string `json:"image"`
}

type OrganisasiResponse struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Jabatan string `json:"jabatan"`
	Nama    string `json:"nama"`
	Image   string `json:"image"`
}
