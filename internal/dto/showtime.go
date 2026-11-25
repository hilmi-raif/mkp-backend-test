package dto

import "time"

type ShowtimeInput struct {
	StudioID int64     `json:"id_studio" validate:"gt=0"`
	FilmID   int64     `json:"id_film" validate:"gt=0"`
	StartsAt time.Time `json:"mulai_at" validate:"required"`
	EndsAt   time.Time `json:"selesai_at" validate:"required"`
	Price    int64     `json:"harga" validate:"gt=0"`
	Status   string    `json:"status" validate:"omitempty,oneof=aktif penjualan_tutup batal"`
}

type Showtime struct {
	ID int64 `json:"id_jadwal"`
	ShowtimeInput
}

type ShowtimePage struct {
	Data []Showtime `json:"data"`
	Meta PageMeta   `json:"meta"`
}
