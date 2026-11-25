package httpapi

import (
	"strconv"

	"mkp-backend-test/internal/dto"
	"mkp-backend-test/internal/helper"
	"mkp-backend-test/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Showtime struct {
	shows    *service.Showtime
	validate *validator.Validate
}

func NewShowtime(shows *service.Showtime, validate *validator.Validate) *Showtime {
	return &Showtime{shows: shows, validate: validate}
}

func (h *Showtime) List(c *fiber.Ctx) error {
	query := c.Queries()
	page, pageErr := helper.QueryInt(query, "page", 1)
	limit, limitErr := helper.QueryInt(query, "limit", 20)
	input := dto.Pagination{Page: page, Limit: limit}
	if pageErr != nil || limitErr != nil || h.validate.Struct(input) != nil {
		return helper.ErrorJSON(c, 400, "page harus bilangan bulat positif maksimal 21474836 dan limit harus 1–100")
	}
	shows, hasNext, err := h.shows.List(c.UserContext(), page, limit)
	if err != nil {
		return helper.HandleError(c, err)
	}
	meta := dto.PageMeta{Page: page, Limit: limit, HasNext: hasNext}
	return c.JSON(dto.ShowtimePage{Data: shows, Meta: meta})
}

func (h *Showtime) Find(c *fiber.Ctx) error {
	id, ok := helper.PositiveID(c.Params("id"))
	if !ok {
		return helper.ErrorJSON(c, 400, "id jadwal tidak valid")
	}
	show, err := h.shows.Find(c.UserContext(), id)
	if err != nil {
		return helper.HandleError(c, err)
	}
	return c.JSON(show)
}

func (h *Showtime) Create(c *fiber.Ctx) error { return h.save(c, 0) }

func (h *Showtime) Update(c *fiber.Ctx) error {
	id, ok := helper.PositiveID(c.Params("id"))
	if !ok {
		return helper.ErrorJSON(c, 400, "id jadwal tidak valid")
	}
	return h.save(c, id)
}

func (h *Showtime) Delete(c *fiber.Ctx) error {
	id, ok := helper.PositiveID(c.Params("id"))
	if !ok {
		return helper.ErrorJSON(c, 400, "id jadwal tidak valid")
	}
	if err := h.shows.Delete(c.UserContext(), id); err != nil {
		return helper.HandleError(c, err)
	}
	return c.Status(204).Send(nil)
}

func (h *Showtime) save(c *fiber.Ctx, id int64) error {
	var input dto.ShowtimeInput
	if c.BodyParser(&input) != nil || h.validate.Struct(input) != nil {
		return helper.ErrorJSON(c, 400, "data jadwal tidak valid")
	}
	if id == 0 {
		show, err := h.shows.Create(c.UserContext(), input)
		if err != nil {
			return helper.HandleError(c, err)
		}
		c.Set(fiber.HeaderLocation, "/jadwal-tayang/"+strconv.FormatInt(show.ID, 10))
		return c.Status(201).JSON(show)
	}
	if input.Status == "" {
		return helper.ErrorJSON(c, 400, "status wajib diisi saat mengubah jadwal")
	}
	show, err := h.shows.Update(c.UserContext(), id, input)
	if err != nil {
		return helper.HandleError(c, err)
	}
	return c.JSON(show)
}
