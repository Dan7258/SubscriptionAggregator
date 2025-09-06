package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// @Description Custom type representing a month and year in MM-YYYY format (e.g., "03-2025").
type MonthYear struct {
	time.Time
}

const monthYearFormat = "01-2006"

func (my *MonthYear) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	if str == "" {
		my.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(monthYearFormat, str)
	if err != nil {
		return err
	}

	my.Time = t
	return nil
}

func (my MonthYear) MarshalJSON() ([]byte, error) {
	if my.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", my.Format(monthYearFormat))), nil
}

func (my MonthYear) Value() (driver.Value, error) {
	if my.Time.IsZero() {
		return nil, nil
	}
	return my.Format("2006-01-02"), nil
}

func (my *MonthYear) Scan(value interface{}) error {
	if value == nil {
		my.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		my.Time = v
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		my.Time = t
	default:
		return fmt.Errorf("cannot scan type %T into MonthYear", value)
	}

	return nil
}
