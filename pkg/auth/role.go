package auth

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Role int

const (
	NilRole Role = iota
	DirectorRole
	AdminRole
	StaffRole
)

var jsonNull = []byte("null")

func ParseUserType(s string) Role {
	switch strings.ToUpper(s) {
	case "DIRECTOR":
		return DirectorRole
	case "ADMIN":
		return AdminRole
	case "STAFF":
		return StaffRole
	default:
		return NilRole
	}
}

func (ut Role) String() string {
	switch ut {
	case DirectorRole:
		return "DIRECTOR"
	case AdminRole:
		return "ADMIN"
	case StaffRole:
		return "STAFF"
	default:
		return "null"
	}
}

func (ut Role) MarshalJSON() ([]byte, error) {
	if ut == NilRole {
		return jsonNull, nil
	}
	str := fmt.Sprintf("\"%s\"", ut)
	return []byte(str), nil
}

func (ut *Role) UnmarshalJSON(raw []byte) error {
	str := string(raw)
	parsed := ParseUserType(strings.Trim(str, `"`))

	if parsed != NilRole {
		*ut = parsed
	}

	return nil
}

func (ut Role) Value() (driver.Value, error) {
	if ut == NilRole {
		return nil, nil
	}
	return ut.String(), nil
}

func (ut *Role) Scan(src any) error {
	if src == nil {
		*ut = NilRole
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan Role from %v", src)
	}

	*ut = ParseUserType(str)

	return nil
}
