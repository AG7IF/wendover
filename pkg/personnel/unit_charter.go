package personnel

import (
	"fmt"
)

type HomeUnit struct {
	region        string
	wing          string
	charterNumber uint
	unitName      string
}

func (h HomeUnit) Region() string {
	return h.region
}

func (h HomeUnit) Wing() string {
	return h.wing
}

func (h HomeUnit) CharterNumber() uint {
	return h.charterNumber
}

func (h HomeUnit) UnitName() string {
	return h.unitName
}

func (h HomeUnit) FullCharterNumber() string {
	return fmt.Sprintf("%s-%s-%3d", h.region, h.wing, h.charterNumber)
}

func (h HomeUnit) ShortCharterNumber() string {
	return fmt.Sprintf("%s-%3d", h.wing, h.charterNumber)
}
