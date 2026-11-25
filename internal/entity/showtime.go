package entity

import "time"

type Showtime struct {
	ID       int64     `gorm:"column:id_jadwal;primaryKey"`
	StudioID int64     `gorm:"column:id_studio"`
	FilmID   int64     `gorm:"column:id_film"`
	StartsAt time.Time `gorm:"column:mulai_at"`
	EndsAt   time.Time `gorm:"column:selesai_at"`
	Price    int64     `gorm:"column:harga"`
	Status   string
}

func (Showtime) TableName() string { return "jadwal_tayang" }
