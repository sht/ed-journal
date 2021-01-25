package event

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Timestamp implements custom json marshaler/unmarshaler interfaces, parsing
// timestamps from RFC 3339 format and encoding into unix timestamps
type Timestamp struct {
	time.Time
}

// Verify json marshaler/unmarshaler interfaces are implemented on compile-time
var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)

// MarshalJSON implements the json.Marshaler interface
func (t Timestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", t.Time.Unix())
	return []byte(stamp), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	var err error
	t.Time, err = time.Parse(time.RFC3339, strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	return nil
}
