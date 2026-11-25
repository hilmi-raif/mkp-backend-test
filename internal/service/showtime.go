package service

import (
	"context"

	"mkp-backend-test/internal/dto"
	"mkp-backend-test/internal/entity"
	"mkp-backend-test/internal/helper"
	"mkp-backend-test/internal/repository"
)

type Showtime struct{ store *repository.Postgres }

func NewShowtime(store *repository.Postgres) *Showtime { return &Showtime{store: store} }

func (s *Showtime) Create(ctx context.Context, input dto.ShowtimeInput) (dto.Showtime, error) {
	if input.Status == "" {
		input.Status = "aktif"
	}
	return s.save(ctx, 0, input)
}

func (s *Showtime) Update(ctx context.Context, id int64, input dto.ShowtimeInput) (dto.Showtime, error) {
	return s.save(ctx, id, input)
}

func (s *Showtime) List(ctx context.Context, page, limit int) ([]dto.Showtime, bool, error) {
	offset := (page - 1) * limit
	shows, err := s.store.List(ctx, limit+1, offset)
	if err != nil {
		return nil, false, err
	}
	hasNext := len(shows) > limit
	if hasNext {
		shows = shows[:limit]
	}
	result := make([]dto.Showtime, 0, len(shows))
	for _, show := range shows {
		result = append(result, helper.ToShowtimeDTO(show))
	}
	return result, hasNext, nil
}

func (s *Showtime) Find(ctx context.Context, id int64) (dto.Showtime, error) {
	show, err := s.store.Find(ctx, id)
	return helper.ToShowtimeDTO(show), err
}

func (s *Showtime) Delete(ctx context.Context, id int64) error { return s.store.Delete(ctx, id) }

func (s *Showtime) save(ctx context.Context, id int64, input dto.ShowtimeInput) (dto.Showtime, error) {
	if !input.EndsAt.After(input.StartsAt) {
		return dto.Showtime{}, helper.ErrTimeRange
	}
	show := entity.Showtime{ID: id, StudioID: input.StudioID, FilmID: input.FilmID, StartsAt: input.StartsAt, EndsAt: input.EndsAt, Price: input.Price, Status: input.Status}
	var err error
	if id == 0 {
		err = s.store.Create(ctx, &show)
	} else {
		err = s.store.Update(ctx, &show)
	}
	return helper.ToShowtimeDTO(show), err
}
