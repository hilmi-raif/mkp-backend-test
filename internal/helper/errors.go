package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	ErrCredentials = errors.New("email atau password salah")
	ErrTimeRange   = errors.New("selesai_at harus setelah mulai_at")
)

func HandleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrCredentials):
		return ErrorJSON(c, 400, err.Error())
	case errors.Is(err, ErrTimeRange):
		return ErrorJSON(c, 400, err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrorJSON(c, 404, "jadwal tidak ditemukan")
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		if c.Method() == fiber.MethodPost {
			return ErrorJSON(c, 400, "studio atau film tidak ditemukan")
		}
		return ErrorJSON(c, 409, "perubahan bertentangan dengan data terkait")
	default:
		return ErrorJSON(c, 500, "terjadi kesalahan pada server")
	}
}
