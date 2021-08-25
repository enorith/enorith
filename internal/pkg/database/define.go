package database

import (
	"fmt"

	"github.com/enorith/supports/carbon"
)

type Datetime struct {
	carbon.Carbon
}

func (c Datetime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.GetDateTimeString())), nil
}

type WithTimestamps struct {
	CreatedAt Datetime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt Datetime `gorm:"column:updated_at" json:"updated_at"`
}
