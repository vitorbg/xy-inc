package model

import (
	"github.com/vitorbg/xy-inc/persistence"
)

func InitDatabase() error {
	return persistence.InitDatabase()
}
