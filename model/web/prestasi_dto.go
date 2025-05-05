package web

type PrestasiCreateRequest struct {
	Title     string `validate:"required"`
	Tahun     string `validate:"required"`
	Prestasi  string `validate:"required"`
	Deskripsi string `validate:"required"`
}

type PrestasiUpdateRequest struct {
	Title     string `validate:"required"`
	Tahun     string `validate:"required"`
	Prestasi  string `validate:"required"`
	Deskripsi string `validate:"required"`
}

type PrestasiResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Tahun     string `json:"tahun"`
	Prestasi  string `json:"prestasi"`
	Deskripsi string `json:"deskripsi"`
}
