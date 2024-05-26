package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

type ShirtSize int

const (
	UnknownSize = iota
	XS
	S
	M
	L
	XL
	XXL
	XXXL
	XXXXL
	XXXXXL
)

func ParseShirtSize(s string) ShirtSize {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "XS":
		return XS
	case "S":
		return S
	case "M":
		return M
	case "L":
		return L
	case "XL":
		return XL
	case "XXL":
		fallthrough
	case "2XL":
		return XXL
	case "XXXL":
		fallthrough
	case "3XL":
		return XXXL
	case "XXXXL":
		fallthrough
	case "4XL":
		return XXXXL
	case "XXXXXL":
		fallthrough
	case "5XL":
		return XXXXXL
	default:
		return UnknownSize
	}
}

func (ss ShirtSize) String() string {
	switch ss {
	case XS:
		return "XS"
	case S:
		return "S"
	case M:
		return "M"
	case L:
		return "L"
	case XL:
		return "XL"
	case XXL:
		return "2XL"
	case XXXL:
		return "3XL"
	case XXXXL:
		return "4XL"
	case XXXXXL:
		return "5XL"
	default:
		panic(errors.Errorf("invalid shirt size: %d", ss))
	}
}
