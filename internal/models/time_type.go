package models

import (
	"fmt"
	"time"
)

type MonthYear struct {
	time.Time
}

const (
	monthYearFormat = "01-2006"
)

func (my MonthYear) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", my.Format(monthYearFormat))
	return []byte(formatted), nil
}

func (my *MonthYear) UnmarshalJSON(b []byte) error {
	str := string(b)
	if len(str) < 2 {
		return fmt.Errorf("invalid date")
	}
	str = str[1 : len(str)-1]

	t, err := time.Parse(monthYearFormat, str)
	if err != nil {
		return err
	}
	my.Time = t
	return nil
}
