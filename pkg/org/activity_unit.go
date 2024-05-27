package org

import (
	"github.com/google/uuid"
)

type ActivityUnit struct {
	id               uuid.UUID
	unitName         string
	subordinateUnits []ActivityUnit
}

func NewActivityUnit(
	id uuid.UUID,
	unitName string,
) ActivityUnit {
	return ActivityUnit{
		id:       id,
		unitName: unitName,
	}
}

func (au *ActivityUnit) ID() uuid.UUID {
	return au.id
}

func (au *ActivityUnit) UnitName() string {
	return au.unitName
}

func (au *ActivityUnit) SubordinateUnits() []ActivityUnit {
	return au.subordinateUnits
}

func (au *ActivityUnit) AddSubordinateUnit(unit ActivityUnit) {
	au.subordinateUnits = append(au.subordinateUnits, unit)
}
