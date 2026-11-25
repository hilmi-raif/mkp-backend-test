package helper

import "gorm.io/gorm"

func CheckRowsAffected(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
