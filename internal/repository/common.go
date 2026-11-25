package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{ db *gorm.DB }

func Open(dsn string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}
