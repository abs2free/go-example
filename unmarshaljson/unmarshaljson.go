package unmarshaljson

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string
	Born Date      `json:"birthdate"`
	Size ShirtSize `json:"shirt-size"`
}

type ShirtSize byte

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
)

var (
	_ShirtSizeNameToValue = map[string]ShirtSize{
		"NA": NA,
		"XS": XS,
		"S":  S,
		"M":  M,
		"L":  L,
		"XL": XL,
	}

	_ShirtSizeValueToName = map[ShirtSize]string{
		NA: "NA",
		XS: "XS",
		S:  "S",
		M:  "M",
		L:  "L",
		XL: "XL",
	}
)

type Date struct{ time.Time }

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("birthdate should be a string, got %s", data)
	}
	t, err := time.Parse("2006/01/02", s)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	d.Time = t
	return nil
}
func (d Date) String() string {
	return d.Format("2006/01/02")
}

func (d Date) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(d).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	return json.Marshal(d)
}

// MarshalJSON is generated so ShirtSize satisfies json.Marshaler.
func (r ShirtSize) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ShirtSizeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ShirtSize: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ShirtSize satisfies json.Unmarshaler.
func (r *ShirtSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ShirtSize should be a string, got %s", data)
	}
	v, ok := _ShirtSizeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ShirtSize %q", s)
	}
	*r = v
	return nil
}
