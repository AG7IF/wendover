package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

type Wing int

const (
	ALWG Wing = 1 + iota
	AZWG
	ARWG
	CAWG
	COWG
	CTWG
	DEWG
	FLWG
	GAWG
	IDWG
	ILWG
	INWG
	IAWG
	KSWG
	KYWG
	LAWG
	MEWG
	MDWG
	MAWG
	MIWG
	MNWG
	MSWG
	MOWG
	MTWG
	DCWG
	NEWG
	NVWG
	NHWG
	NJWG
	NMWG
	NYWG
	NCWG
	NDWG
	OHWG
	OKWG
	ORWG
	PAWG
	RIWG
	SCWG
	SDWG
	TNWG
	TXWG
	UTWG
	VTWG
	VAWG
	WAWG
	WVWG
	WIWG
	WYWG
	AKWG
	HIWG
	PRWG
	NER = 39 + iota
	MAR
	GLR
	SER
	NCR
	SWR
	RMR
	PCR
	NHQ
)

func ParseWing(s string) (Wing, error) {
	switch strings.ToUpper(s) {
	case "AL":
		return ALWG, nil
	case "AZ":
		return AZWG, nil
	case "AR":
		return ARWG, nil
	case "CA":
		return CAWG, nil
	case "CO":
		return COWG, nil
	case "CT":
		return CTWG, nil
	case "DE":
		return DEWG, nil
	case "FL":
		return FLWG, nil
	case "GA":
		return GAWG, nil
	case "ID":
		return IDWG, nil
	case "IL":
		return ILWG, nil
	case "IN":
		return INWG, nil
	case "IA":
		return IAWG, nil
	case "KS":
		return KSWG, nil
	case "KY":
		return KYWG, nil
	case "LA":
		return LAWG, nil
	case "ME":
		return MEWG, nil
	case "MD":
		return MDWG, nil
	case "MA":
		return MAWG, nil
	case "MI":
		return MIWG, nil
	case "MN":
		return MNWG, nil
	case "MS":
		return MSWG, nil
	case "MO":
		return MOWG, nil
	case "MT":
		return MTWG, nil
	case "DC":
		return DCWG, nil
	case "NE":
		return NEWG, nil
	case "NV":
		return NVWG, nil
	case "NH":
		return NHWG, nil
	case "NJ":
		return NJWG, nil
	case "NM":
		return NMWG, nil
	case "NY":
		return NYWG, nil
	case "NC":
		return NCWG, nil
	case "ND":
		return NDWG, nil
	case "OH":
		return OHWG, nil
	case "OK":
		return OKWG, nil
	case "OR":
		return ORWG, nil
	case "PA":
		return PAWG, nil
	case "RI":
		return RIWG, nil
	case "SC":
		return SCWG, nil
	case "SD":
		return SDWG, nil
	case "TN":
		return TNWG, nil
	case "TX":
		return TXWG, nil
	case "UT":
		return UTWG, nil
	case "VT":
		return VTWG, nil
	case "VA":
		return VAWG, nil
	case "WA":
		return WAWG, nil
	case "WV":
		return WVWG, nil
	case "WI":
		return WIWG, nil
	case "WY":
		return WYWG, nil
	case "AK":
		return AKWG, nil
	case "HI":
		return HIWG, nil
	case "PR":
		return PRWG, nil
	case "MAR":
		return MAR, nil
	case "GLR":
		return GLR, nil
	case "SER":
		return SER, nil
	case "NCR":
		return NCR, nil
	case "SWR":
		return SWR, nil
	case "RMR":
		return RMR, nil
	case "PCR":
		return PCR, nil
	case "NHQ":
		return NHQ, nil
	default:
		return -1, errors.Errorf("invalid wing: %s", s)
	}
}

func (w Wing) String() string {
	switch w {
	case ALWG:
		return "AL"
	case AZWG:
		return "AZ"
	case ARWG:
		return "AR"
	case CAWG:
		return "CA"
	case COWG:
		return "CO"
	case CTWG:
		return "CT"
	case DEWG:
		return "DE"
	case FLWG:
		return "FL"
	case GAWG:
		return "GA"
	case IDWG:
		return "ID"
	case ILWG:
		return "IL"
	case INWG:
		return "IN"
	case IAWG:
		return "IA"
	case KSWG:
		return "KS"
	case KYWG:
		return "KY"
	case LAWG:
		return "LA"
	case MEWG:
		return "ME"
	case MDWG:
		return "MD"
	case MAWG:
		return "MA"
	case MIWG:
		return "MI"
	case MNWG:
		return "MN"
	case MSWG:
		return "MS"
	case MOWG:
		return "MO"
	case MTWG:
		return "MT"
	case DCWG:
		return "DC"
	case NEWG:
		return "NE"
	case NVWG:
		return "NV"
	case NHWG:
		return "NH"
	case NJWG:
		return "NJ"
	case NMWG:
		return "NM"
	case NYWG:
		return "NY"
	case NCWG:
		return "NC"
	case NDWG:
		return "ND"
	case OHWG:
		return "OH"
	case OKWG:
		return "OK"
	case ORWG:
		return "OR"
	case PAWG:
		return "PA"
	case RIWG:
		return "RI"
	case SCWG:
		return "SC"
	case SDWG:
		return "SD"
	case TNWG:
		return "TN"
	case TXWG:
		return "TX"
	case UTWG:
		return "UT"
	case VTWG:
		return "VT"
	case VAWG:
		return "VA"
	case WAWG:
		return "WA"
	case WVWG:
		return "WV"
	case WIWG:
		return "WI"
	case WYWG:
		return "WY"
	case AKWG:
		return "AK"
	case HIWG:
		return "HI"
	case PRWG:
		return "PR"
	case NER:
		return "NER"
	case MAR:
		return "MAR"
	case GLR:
		return "GLR"
	case SER:
		return "SER"
	case NCR:
		return "NCR"
	case SWR:
		return "SWR"
	case RMR:
		return "RMR"
	case PCR:
		return "PCR"
	case NHQ:
		return "NH"
	default:
		panic(errors.Errorf("invalid wing value: %d", w))
	}
}
