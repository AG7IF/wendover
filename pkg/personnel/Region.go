package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

var wingsByRegion = map[Region]map[Wing]bool{
	NorthEastRegion: {
		NER:  true,
		CTWG: true,
		MAWG: true,
		MEWG: true,
		NHWG: true,
		NJWG: true,
		NYWG: true,
		PAWG: true,
		RIWG: true,
		VTWG: true,
	},
	MidAtlanticRegion: {
		MAR:  true,
		DCWG: true,
		DEWG: true,
		MDWG: true,
		NCWG: true,
		SCWG: true,
		VAWG: true,
		WVWG: true,
	},
	GreatLakesRegion: {
		GLR:  true,
		ILWG: true,
		INWG: true,
		KYWG: true,
		MIWG: true,
		OHWG: true,
		WIWG: true,
	},
	SoutheastRegion: {
		SER:  true,
		ALWG: true,
		FLWG: true,
		GAWG: true,
		MSWG: true,
		PRWG: true,
		TNWG: true,
	},
	NorthCentralRegion: {
		NCR:  true,
		IAWG: true,
		KSWG: true,
		MNWG: true,
		MOWG: true,
		NDWG: true,
		NEWG: true,
		SDWG: true,
	},
	SouthwestRegion: {
		SWR:  true,
		ARWG: true,
		AZWG: true,
		LAWG: true,
		NMWG: true,
		OKWG: true,
		TXWG: true,
	},
	RockyMountainRegion: {
		RMR:  true,
		COWG: true,
		IDWG: true,
		MTWG: true,
		UTWG: true,
		WYWG: true,
	},
	PacificRegion: {
		PCR:  true,
		AKWG: true,
		CAWG: true,
		HIWG: true,
		NVWG: true,
		ORWG: true,
		WAWG: true,
	},
	NationalHeadquarters: {
		NHQ: true,
	},
}

type Region int

const (
	NorthEastRegion Region = 91 + iota
	MidAtlanticRegion
	GreatLakesRegion
	SoutheastRegion
	NorthCentralRegion
	SouthwestRegion
	RockyMountainRegion
	PacificRegion
	NationalHeadquarters
)

func ParseRegion(s string) (Region, error) {
	switch strings.ToUpper(s) {
	case "NER":
		return NorthEastRegion, nil
	case "MAR":
		return MidAtlanticRegion, nil
	case "GLR":
		return GreatLakesRegion, nil
	case "SER":
		return SoutheastRegion, nil
	case "NCR":
		return NorthCentralRegion, nil
	case "SWR":
		return SouthwestRegion, nil
	case "RMR":
		return RockyMountainRegion, nil
	case "PCR":
		return PacificRegion, nil
	case "NHQ":
		return NationalHeadquarters, nil
	default:
		return -1, errors.Errorf("invalid region: %s", s)
	}
}

func (r Region) String() string {
	switch r {
	case NorthEastRegion:
		return "NER"
	case MidAtlanticRegion:
		return "MAR"
	case GreatLakesRegion:
		return "GLR"
	case SoutheastRegion:
		return "SER"
	case NorthCentralRegion:
		return "NCR"
	case SouthwestRegion:
		return "SWR"
	case RockyMountainRegion:
		return "RMR"
	case PacificRegion:
		return "PCR"
	case NationalHeadquarters:
		return "NHQ"
	default:
		panic(errors.Errorf("invalid region value: %d", r))
	}
}

func (r Region) WingIsInRegion(w Wing) bool {
	_, ok := wingsByRegion[r][w]
	return ok
}
