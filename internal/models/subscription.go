package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	ServiceName string    `json:"service_name"`
	Price       uint      `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
}
