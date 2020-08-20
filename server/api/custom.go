package api

import (
	"encoding/json"
	"time"
)

type CustomDate struct {
	time time.Time
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	// your unmarshal logic...
	d.time = time.Now()
	return nil
}

func (d CustomDate) MarshalJSON() ([]byte, error) {
	// your marshal logic...
	return json.Marshal(d.time.Format(time.ANSIC))
}
