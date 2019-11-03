package main

import (
	"github.com/chasdevs/meetrics/pkg/setup"
)

func main() {
	setup.Setup()
	setup.Migrate()
	setup.PopulateUsersFromCsv()
}
