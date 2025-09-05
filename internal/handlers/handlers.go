package handlers

import (
	"subagg/internal/models"
)

type Handlers struct {
	db *models.PostgresDatabase
}

func NewHandlers(db *models.PostgresDatabase) *Handlers {
	return &Handlers{db}
}
