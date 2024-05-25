package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

type Gender int

const (
	NonBinary Gender = iota
	Male
	Female
)

func ParseGender(s string) Gender {
	switch strings.ToLower(s) {
	case "m":
		fallthrough
	case "male":
		return Male
	case "f":
		fallthrough
	case "female":
		return Female
	default:
		return NonBinary
	}
}

func (g Gender) String() string {
	switch g {
	case NonBinary:
		return "Non-Binary"
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		panic(errors.Errorf("invalid gender: %d", g))
	}
}
