package main

import (
	"log"

	"github.com/orca-cpfr/orca-cpfr.github.io/landing-page/internal/app"
)

func main() {
	err := app.Render()
	if err != nil {
		log.Fatal(err)
	}
}
