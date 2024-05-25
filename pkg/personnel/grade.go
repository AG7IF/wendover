package personnel

import (
	"github.com/pkg/errors"
)

type Grade int

const (
	CdtAB Grade = iota
	CdtAmn
	CdtA1C
	CdtSrA
	CdtSSgt
	CdtTSgt
	CdtMSgt
	CdtSMSgt
	CdtCMSgt
	CdtSecondLt
	CdtFirstLt
	CdtCapt
	CdtMaj
	CdtLtCol
	CdtCol
	SM
	SSgt
	TSgt
	MSgt
	SMSgt
	CMSgt
	FO
	TFO
	SFO
	SecondLt
	FirstLt
	Capt
	Maj
	LtCol
	Col
	BrigGen
	MajGen
)

func ParseGrade(s string) (Grade, error) {
	switch s {
	case "Maj Gen":
		return MajGen, nil
	case "Brig Gen":
		return BrigGen, nil
	case "Col":
		return Col, nil
	case "Lt Col":
		return LtCol, nil
	case "Maj":
		return Maj, nil
	case "Capt":
		return Capt, nil
	case "1st Lt":
		return FirstLt, nil
	case "2d Lt":
		return SecondLt, nil
	case "SFO":
		return SFO, nil
	case "TFO":
		return TFO, nil
	case "FO":
		return FO, nil
	case "CMSgt":
		return CMSgt, nil
	case "SMSgt":
		return SMSgt, nil
	case "MSgt":
		return MSgt, nil
	case "TSgt":
		return TSgt, nil
	case "SSgt":
		return SSgt, nil
	case "SM":
		return SM, nil
	case "C/Col":
		return CdtCol, nil
	case "C/Lt Col":
		return CdtLtCol, nil
	case "C/Maj":
		return CdtMaj, nil
	case "C/Capt":
		return CdtCapt, nil
	case "C/1st Lt":
		return CdtFirstLt, nil
	case "C/2d Lt":
		return CdtSecondLt, nil
	case "C/CMSgt":
		return CdtCMSgt, nil
	case "C/SMSgt":
		return CdtSMSgt, nil
	case "C/MSgt":
		return CdtMSgt, nil
	case "C/TSgt":
		return CdtTSgt, nil
	case "C/SSgt":
		return CdtSSgt, nil
	case "C/SrA":
		return CdtSrA, nil
	case "C/A1C":
		return CdtA1C, nil
	case "C/Amn":
		return CdtAmn, nil
	case "C/AB":
		fallthrough
	case "CADET":
		return CdtAB, nil
	default:
		return -1, errors.Errorf("invalid member grade: %s", s)
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
	case SFO:
		return "SFO"
	case TFO:
		return "TFO"
	case FO:
		return "FO"
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
		panic(errors.Errorf("invalid grade: %d", g))
	}
}
