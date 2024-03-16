package unmarshaljson

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type PersonOld struct {
	Name string
	Born time.Time
	Size ShirtSize
}

func (p *PersonOld) Parse(s string) error {
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&fields); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	// Once decoded we can access the fields by name.
	p.Name = fields["name"]

	born, err := time.Parse("2006/01/02", fields["birthdate"])
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	p.Born = born

	p.Size, err = ParseShirtSize(fields["shirt-size"])
	return err
}

func ParseShirtSize(s string) (ShirtSize, error) {
	sizes := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}
	ss, ok := sizes[s]
	if !ok {
		return NA, fmt.Errorf("invalid ShirtSize %q", s)
	}
	return ss, nil
}
