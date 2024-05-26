package personnel

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

type HomeUnit struct {
	region     Region
	wing       Wing
	unitNumber uint
	unitName   string
}

func NewHomeUnit(region Region, wing Wing, unitNumber uint, unitName string) (HomeUnit, error) {
	if !region.WingIsInRegion(wing) {
		return HomeUnit{}, errors.Errorf("%s is not a part of %s", wing, region)
	}

	return HomeUnit{
		region:     region,
		wing:       wing,
		unitNumber: unitNumber,
		unitName:   unitName,
	}, nil
}

func ParseHomeUnit(charter string, unitName string) (HomeUnit, error) {
	r := regexp.MustCompile(`(\w{3})-(\w{2,3})-(\d{3})`)

	m := r.FindStringSubmatch(charter)

	if len(m) < 4 {
		return HomeUnit{}, errors.Errorf("malformed unit charter number: %s", charter)
	}

	region, err := ParseRegion(m[1])
	if err != nil {
		return HomeUnit{}, errors.WithMessagef(err, "unable to parse region from charter number: %s", charter)
	}

	wing, err := ParseWing(m[2])
	if err != nil {
		return HomeUnit{}, errors.WithMessagef(err, "unable to parse wing from charter number: %s", charter)
	}

	unitNumber, err := strconv.Atoi(m[3])
	if err != nil {
		return HomeUnit{}, errors.WithMessagef(err, "unable to parse unit number from charter number: %s", charter)
	}

	if !region.WingIsInRegion(wing) {
		return HomeUnit{}, errors.Errorf("%s is not a part of %s", wing, region)
	}

	return HomeUnit{
		region:     region,
		wing:       wing,
		unitNumber: uint(unitNumber),
		unitName:   unitName,
	}, nil
}

func (h HomeUnit) Region() Region {
	return h.region
}

func (h HomeUnit) Wing() Wing {
	return h.wing
}

func (h HomeUnit) UnitNumber() uint {
	return h.unitNumber
}

func (h HomeUnit) UnitName() string {
	return h.unitName
}

func (h HomeUnit) ShortCharterNumber() string {
	return fmt.Sprintf("%s-%03d", h.wing, h.unitNumber)
}

func (h HomeUnit) FullCharterNumber() string {
	return fmt.Sprintf("%s-%s-%03d", h.region, h.wing, h.unitNumber)
}
