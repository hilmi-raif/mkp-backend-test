package entity

type User struct {
	ID           int64 `gorm:"column:id_pengguna;primaryKey"`
	Email        string
	PasswordHash string
}

func (User) TableName() string { return "pengguna" }
