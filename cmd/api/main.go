package main

import (
	"log"
	"mkp-backend-test/internal/bootstrap"
	"mkp-backend-test/internal/config"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	app, err := bootstrap.New(c)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Listen(c.Addr); err != nil {
		log.Fatal(err)
	}
}
