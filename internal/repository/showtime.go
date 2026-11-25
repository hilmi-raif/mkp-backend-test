package repository

import (
	"context"

	"mkp-backend-test/internal/entity"

	"mkp-backend-test/internal/helper"
)

func (r *Postgres) Create(ctx context.Context, show *entity.Showtime) error {
	return r.db.WithContext(ctx).Create(show).Error
}

func (r *Postgres) List(ctx context.Context, limit, offset int) ([]entity.Showtime, error) {
	shows := []entity.Showtime{}
	err := r.db.WithContext(ctx).Order("id_jadwal").Limit(limit).Offset(offset).Find(&shows).Error
	return shows, err
}

func (r *Postgres) Find(ctx context.Context, id int64) (entity.Showtime, error) {
	var show entity.Showtime
	err := r.db.WithContext(ctx).Take(&show, "id_jadwal = ?", id).Error
	return show, err
}

func (r *Postgres) Update(ctx context.Context, show *entity.Showtime) error {
	db := r.db.WithContext(ctx).Model(show).Select("id_studio", "id_film", "mulai_at", "selesai_at", "harga", "status").Updates(show)
	return helper.CheckRowsAffected(db)
}

func (r *Postgres) Delete(ctx context.Context, id int64) error {
	return helper.CheckRowsAffected(r.db.WithContext(ctx).Delete(&entity.Showtime{}, "id_jadwal = ?", id))
}
