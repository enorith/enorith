package database

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/enorith/supports/carbon"
)

type Datetime struct {
	carbon.Carbon
}

func (c Datetime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.GetDateTimeString())), nil
}

type WithTimestamps struct {
	CreatedAt Datetime  `gorm:"column:created_at;autoCreateTime;type:timestamp null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamp null" json:"updated_at"`
}

type SliceString []string

// Scan assigns a value from a database driver.
// The src value will be of one of the following types:
//
//	int64
//	float64
//	bool
//	[]byte
//	string
//	time.Time
//	nil - for NULL values
//
// An error should be returned if the value cannot be stored
// without loss of information.
//
// Reference types such as []byte are only valid until the next call to Scan
// and should not be retained. Their underlying memory is owned by the driver.
// If retention is necessary, copy their values before the next call to Scan.
func (ss *SliceString) Scan(src any) error {
	var val string
	if s, ok := src.(string); ok {
		val = s
	}

	if s, ok := src.([]byte); ok {
		val = string(s)
	}
	if val == "" {
		*ss = make(SliceString, 0)
	} else {
		*ss = strings.Split(val, ",")
	}

	return nil
}

func (ss SliceString) Value() (driver.Value, error) {
	return strings.Join(ss, ","), nil
}
