package personnel

import (
	"strings"

	"github.com/pkg/errors"
)

type ShirtSize int

const (
	XXXS ShirtSize = iota
	XXS
	XS
	S
	M
	L
	XL
	XXL
	XXXL
)

func ParseShirtSize(s string) (ShirtSize, error) {
	switch strings.ToLower(s) {
	case "xxxs":
		return XXXS, nil
	case "xxs":
		return XXS, nil
	case "xs":
		return XS, nil
	case "s":
		return S, nil
	case "m":
		return M, nil
	case "l":
		return L, nil
	case "xl":
		return XL, nil
	case "xxl":
		return XXL, nil
	case "xxxl":
		return XXXL, nil
	default:
		return -1, errors.Errorf("unrecognized shirt size: %s", s)
	}
}

func (ss ShirtSize) String() string {
	switch ss {
	case XXXS:
		return "XXXS"
	case XXS:
		return "XXS"
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
		return "XXL"
	case XXXL:
		return "XXXL"
	default:
		panic(errors.Errorf("invalid shirt size: %d", ss))
	}
}
