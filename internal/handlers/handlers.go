package handlers

import (
	"github.com/gin-gonic/gin"
	"subagg/internal/models"
)

type Handlers struct {
	db *models.PostgresDatabase
}

func NewSubscriptionHandlers(db *models.PostgresDatabase) *Handlers {
	return &Handlers{db}
}
