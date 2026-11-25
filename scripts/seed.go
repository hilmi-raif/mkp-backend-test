package main

import (
	"fmt"
	"log"

	"mkp-backend-test/internal/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn, err := config.DatabaseURL()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("mkp-demo-password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	var branch, studio, film int64
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`insert into pengguna (email, password_hash) values (?, ?) on conflict (email) do nothing`, "demo@example.com", string(hash)).Error; err != nil {
			return err
		}
		if err := tx.Raw(`insert into cabang (nama, kota) values ('Bioskop Demo', 'Semarang') on conflict (kota, nama) do update set nama = excluded.nama returning id_cabang`).Scan(&branch).Error; err != nil {
			return err
		}
		if err := tx.Raw(`insert into studio (id_cabang, nama) values (?, 'Studio Demo') on conflict (id_cabang, nama) do update set nama = excluded.nama returning id_studio`, branch).Scan(&studio).Error; err != nil {
			return err
		}
		if err := tx.Raw(`select id_film from film where judul = 'Film Demo' order by id_film limit 1`).Scan(&film).Error; err != nil {
			return err
		}
		if film == 0 {
			return tx.Raw(`insert into film (judul) values ('Film Demo') returning id_film`).Scan(&film).Error
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Demo: demo@example.com | id_studio=%d | id_film=%d\n", studio, film)
}
