package setup

import (
	"github.com/chasdevs/meetrics/pkg/data"
)

func Migrate() {
	data.Migrate()
}

func Setup() {
	data.SetupDb()
}

func TearDown() {
	data.TeardownDb()
}
