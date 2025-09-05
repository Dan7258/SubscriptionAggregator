package models

import "github.com/google/uuid"

type SubscriptionFilters struct {
	UserID    uuid.UUID `json:"user_id"`
	StartDate MonthYear `json:"start_date"`
	EndDate   MonthYear `json:"end_date"`
}

func (db PostgresDatabase) GetSubscriptionsByFilters(filters SubscriptionFilters) ([]Subscription, error) {
	var subscriptions []Subscription
	query := db.Conn.Model(&subscriptions)
	if filters.UserID != uuid.Nil {
		query = query.Where("user_id = ?", filters.UserID).Find(&subscriptions)
	}
	if !filters.StartDate.IsZero() {
		query = query.Where("start_date >= ?", filters.StartDate)
	}
	if !filters.EndDate.IsZero() {
		query = query.Where("start_date <= ?", filters.EndDate)
	}
	err := query.Find(&subscriptions).Error
	return subscriptions, err
}
