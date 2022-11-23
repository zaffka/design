package domain

import (
	"strings"
	"time"
)

const CustomTimeLayout = "2006-01-02"

type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses time from the b using a CustomTimeLayout.
// CustomTime gets a zero Time if any parsing errors occurs.
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	ct.Time, err = time.Parse(CustomTimeLayout, strings.Trim(string(b), "\""))
	if err != nil {
		ct.Time = time.Time{}

		return nil
	}

	return nil
}
