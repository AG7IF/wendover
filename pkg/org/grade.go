package org

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/ag7if/wendover/pkg"
)

type Grade int

const (
	NilGrade Grade = iota
	MajGen
	BrigGen
	Col
	LtCol
	Maj
	Capt
	FirstLt
	SecondLt
	CMSgt
	SMSgt
	MSgt
	TSgt
	SSgt
	SFO
	TFO
	FO
	SM
	CdtCol
	CdtLtCol
	CdtMaj
	CdtCapt
	CdtFirstLt
	CdtSecondLt
	CdtCMSgt
	CdtSMSgt
	CdtMSgt
	CdtTSgt
	CdtSSgt
	CdtSrA
	CdtA1C
	CdtAmn
	CdtAB
)

func ParseGrade(s string) Grade {
	switch s {
	case "Maj Gen":
		return MajGen
	case "Brig Gen":
		return BrigGen
	case "Col":
		return Col
	case "Lt Col":
		return LtCol
	case "Maj":
		return Maj
	case "Capt":
		return Capt
	case "1st Lt":
		return FirstLt
	case "2d Lt":
		return SecondLt
	case "CMSgt":
		return CMSgt
	case "SMSgt":
		return SMSgt
	case "TSgt":
		return TSgt
	case "SSgt":
		return SSgt
	case "SFO":
		return SFO
	case "TFO":
		return TFO
	case "FO":
		return FO
	case "SM":
		fallthrough
	case "SENIOR":
		return SM
	case "C/Col":
		return CdtCol
	case "C/Lt Col":
		return CdtLtCol
	case "C/Maj":
		return CdtMaj
	case "C/Capt":
		return CdtCapt
	case "C/1st Lt":
		return CdtFirstLt
	case "C/2d Lt":
		return CdtSecondLt
	case "C/CMSgt":
		return CdtCMSgt
	case "C/SMSgt":
		return CdtSMSgt
	case "C/MSgt":
		return CdtMSgt
	case "C/TSgt":
		return CdtTSgt
	case "C/SSgt":
		return CdtSSgt
	case "C/SrA":
		return CdtSrA
	case "C/A1C":
		return CdtA1C
	case "C/Amn":
		return CdtAmn
	case "C/AB":
		fallthrough
	case "CADET":
		return CdtAB
	default:
		return NilGrade
	}
}

func (g Grade) String() string {
	switch g {
	case MajGen:
		return "Maj Gen"
	case BrigGen:
		return "Brig Gen"
	case Col:
		return "Col"
	case LtCol:
		return "Lt Col"
	case Maj:
		return "Maj"
	case Capt:
		return "Capt"
	case FirstLt:
		return "1st Lt"
	case SecondLt:
		return "2d Lt"
	case CMSgt:
		return "CMSgt"
	case SMSgt:
		return "SMSgt"
	case MSgt:
		return "MSgt"
	case TSgt:
		return "TSgt"
	case SSgt:
		return "SSgt"
	case SFO:
		return "SFO"
	case TFO:
		return "TFO"
	case FO:
		return "FO"
	case SM:
		return "SM"
	case CdtCol:
		return "C/Col"
	case CdtLtCol:
		return "C/Lt Col"
	case CdtMaj:
		return "C/Maj"
	case CdtCapt:
		return "C/Capt"
	case CdtFirstLt:
		return "C/1st Lt"
	case CdtSecondLt:
		return "C/2d Lt"
	case CdtCMSgt:
		return "C/CMSgt"
	case CdtSMSgt:
		return "C/SMSgt"
	case CdtMSgt:
		return "C/MSgt"
	case CdtTSgt:
		return "C/TSgt"
	case CdtSSgt:
		return "C/SSgt"
	case CdtSrA:
		return "C/SrA"
	case CdtA1C:
		return "C/A1C"
	case CdtAmn:
		return "C/Amn"
	case CdtAB:
		return "C/AB"
	default:
		return ""
	}
}

func (g Grade) MarshalJSON() ([]byte, error) {
	if g == NilGrade {
		return pkg.JSONNull, nil
	}
	str := fmt.Sprintf("\"%s\"", g)
	return []byte(str), nil
}

func (g *Grade) UnmarshalJSON(raw []byte) error {
	str := string(raw)
	parsed := ParseGrade(strings.Trim(str, `""`))

	if parsed != NilGrade {
		*g = parsed
	}

	return nil
}

func (g Grade) Value() (driver.Value, error) {
	if g == NilGrade {
		return nil, nil
	}

	return g.String(), nil
}

func (g *Grade) Scan(src any) error {
	if src == nil {
		*g = NilGrade
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan Grade from %v", src)
	}

	*g = ParseGrade(str)

	return nil
}
