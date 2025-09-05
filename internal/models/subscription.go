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

func (db PostgresDatabase) CreateSubscription(subscription Subscription) error {
	err := db.Conn.Create(&subscription).Error
	return err
}

func (db PostgresDatabase) GetSubscriptionByID(id uint) (Subscription, error) {
	var subscription Subscription
	err := db.Conn.First(&subscription, id).Error
	return subscription, err
}

func (db PostgresDatabase) UpdateSubscriptionByID(subscription Subscription) error {
	err := db.Conn.Save(&subscription).Error
	return err
}

func (db PostgresDatabase) DeleteSubscriptionByID(id uint) error {
	err := db.Conn.Delete(&Subscription{}, id).Error
	return err
}

func (db PostgresDatabase) GetSubscriptions() ([]Subscription, error) {
	var subscriptions []Subscription
	err := db.Conn.Find(&subscriptions).Error
	return subscriptions, err
}
