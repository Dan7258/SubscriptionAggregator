package models

import "github.com/google/uuid"

// @Description Filters for retrieving subscriptions based on specific criteria.
type SubscriptionFilters struct {
	// The unique identifier of the user (UUID format) to filter subscriptions.
	// Required: false
	// Format: uuid
	UserID uuid.UUID `json:"user_id"`

	// The start date to filter subscriptions (in MM-YYYY format, e.g., "03-2025").
	// Required: false
	// Example: 03-2025
	StartDate MonthYear `json:"start_date" swaggertype:"string"`

	// The end date to filter subscriptions (in MM-YYYY format, e.g., "12-2025").
	// Required: false
	// Example: 12-2025
	EndDate MonthYear `json:"end_date" swaggertype:"string"`
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
