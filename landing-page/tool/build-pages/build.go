package main

import (
	"log"

	"github.com/orca-cpfr/orca-cpfr.github.io/landing-page/internal/generator"
)

func main() {
	err := generator.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
