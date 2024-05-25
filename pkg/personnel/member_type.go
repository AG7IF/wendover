package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	PatronMember MemberType = iota
	SeniorMember
	CadetSponsorMember
	CadetMember
)

func ParseMemberType(s string) MemberType {
	switch strings.ToLower(s) {
	case "senior":
		return SeniorMember
	case "cadet sponsor":
		return CadetSponsorMember
	case "cadet":
		return CadetMember
	default:
		return PatronMember
	}
}

func (mt MemberType) String() string {
	switch mt {
	case PatronMember:
		return "PATRON"
	case SeniorMember:
		return "SENIOR"
	case CadetSponsorMember:
		return "CADET SPONSOR"
	case CadetMember:
		return "CADET"
	default:
		panic(errors.Errorf("invalid member type: %d", mt))
	}
}
