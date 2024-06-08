package views

import (
	"github.com/google/uuid"

	"github.com/ag7if/wendover/pkg/org"
)

type ActivityUnitView struct {
	ID               uuid.UUID          `json:"id"`
	UnitName         string             `json:"unit_name"`
	SubordinateUnits []ActivityUnitView `json:"subordinateUnits"`
}

func NewActivityUnitView(unit org.ActivityUnit) ActivityUnitView {
	auv := ActivityUnitView{
		ID:       unit.ID(),
		UnitName: unit.UnitName(),
	}

	subunits := unit.SubordinateUnits()

	if len(subunits) == 0 {
		return auv
	}

	for _, unit := range subunits {
		v := NewActivityUnitView(unit)
		auv.SubordinateUnits = append(auv.SubordinateUnits, v)
	}

	return auv
}

func (auv ActivityUnitView) ToDomainObject() org.ActivityUnit {
	au := org.NewActivityUnit(auv.ID, auv.UnitName)

	if len(auv.SubordinateUnits) == 0 {
		return au
	}

	for _, v := range auv.SubordinateUnits {
		u := v.ToDomainObject()
		au.AddSubordinateUnit(u)
	}

	return au
}
