package models

import (
	"fmt"
	"time"
)

// Date represents a calendar date without time component
type Date struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Accepts date in YYYY-MM-DD format
func (d *Date) UnmarshalJSON(data []byte) error {
	// Remove quotes
	s := string(data)
	s = s[1 : len(s)-1]

	// Parse date in YYYY-MM-DD format
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %v", err)
	}

	d.Time = t
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// Returns date in YYYY-MM-DD format
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

// String returns the date in YYYY-MM-DD format
func (d Date) String() string {
	return d.Format("2006-01-02")
} 