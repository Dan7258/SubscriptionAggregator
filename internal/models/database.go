package models

import (
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Conn *gorm.DB
}

func NewDatabase(gorm *gorm.DB) *PostgresDatabase {
	return &PostgresDatabase{gorm}
}
