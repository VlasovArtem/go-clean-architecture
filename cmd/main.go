package main

import (
	"github.com/rs/zerolog/log"

	"clean-architecture/internal/app"
)

func main() {
	application, err := app.NewApplication()

	if err != nil {
		log.Fatal().Err(err)
	}

	if err = application.Run(); err != nil {
		log.Fatal().Err(err)
	}
}
