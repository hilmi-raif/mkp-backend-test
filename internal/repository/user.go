package repository

import (
	"context"

	"mkp-backend-test/internal/entity"
)

func (r *Postgres) FindUser(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	return user, err
}
