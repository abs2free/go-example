package unmarshaljson

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

var tm, _ = time.Parse("2006/01/02", "2009/11/10")

func TestParse(t *testing.T) {
	var input = `{
		"name": "Gopher",
		"birthdate": "2009/11/10",
		"shirt-size": "XS"
	    }`

	var p PersonOld
	if err := p.Parse(input); err != nil {
		log.Fatalf("parse person: %v", err)
	}

	assert.Equal(t, p, PersonOld{
		Name: "Gopher",
		Born: tm,
		Size: XS,
	})

}

func TestUnmarshal(t *testing.T) {
	var input = `{
		"name": "Gopher",
		"birthdate": "2009/11/10",
		"shirt-size": "XS"
	    }`

	var p Person
	dec := json.NewDecoder(strings.NewReader(input))
	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}

	tm, _ = time.Parse("2006/01/02", "2009/11/10")

	assert.Equal(t, p, Person{
		Name: "Gopher",
		Born: Date{tm},
		Size: XS,
	})
}

func TestMarshal(t *testing.T) {
	p := Person{
		Name: "hello",
		Born: Date{tm},
		Size: XL,
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal("JSON encoding failed: ", err)
	}
	if string(data) != `{"Name":"hello","birthdate":"2009/11/10","shirt-size":"XL"}` {
		log.Fatalf("Json encodeing falild: %s", string(data))
	}
}
