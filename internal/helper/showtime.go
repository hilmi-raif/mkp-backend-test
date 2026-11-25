package helper

import (
	"mkp-backend-test/internal/dto"
	"mkp-backend-test/internal/entity"
)

func ToShowtimeDTO(show entity.Showtime) dto.Showtime {
	return dto.Showtime{ID: show.ID, ShowtimeInput: dto.ShowtimeInput{
		StudioID: show.StudioID, FilmID: show.FilmID, StartsAt: show.StartsAt, EndsAt: show.EndsAt, Price: show.Price, Status: show.Status,
	}}
}
