package web

type BeritaCreateRequest struct {
	TitleBerita string `validate:"required"`
	Tanggal     string `validate:"required"`
	Deskripsi   string `validate:"required"`
}

type BeritaUpdateRequest struct {
	TitleBerita string `validate:"required"`
	Tanggal     string `validate:"required"`
	Deskripsi   string `validate:"required"`
}

type BeritaResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	TitleBerita string `json:"title_berita"`
	Tanggal     string `json:"tanggal"`
	Image       string `json:"image"`
	Deskripsi   string `json:"deskripsi"`
}
