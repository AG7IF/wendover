package org

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	id               uuid.UUID
	key              string
	name             string
	location         string
	start            time.Time
	end              time.Time
	cadetStudentFee  uint
	cadetCadreFee    uint
	seniorStudentFee uint
	seniorCadreFee   uint
}

func NewActivity(
	id uuid.UUID,
	key string,
	name string,
	location string,
	start time.Time,
	end time.Time,
	cadetStudentFee uint,
	cadetCadreFee uint,
	seniorStudentFee uint,
	seniorCadreFee uint,
) Activity {
	return Activity{
		id:               id,
		key:              strings.ToUpper(key),
		name:             name,
		location:         location,
		start:            start,
		end:              end,
		cadetStudentFee:  cadetStudentFee,
		cadetCadreFee:    cadetCadreFee,
		seniorStudentFee: seniorStudentFee,
		seniorCadreFee:   seniorCadreFee,
	}
}

func (a Activity) ID() uuid.UUID {
	return a.id
}

func (a Activity) Key() string {
	return a.key
}

func (a Activity) Name() string {
	return a.name
}

func (a Activity) Location() string {
	return a.location
}

func (a Activity) Start() time.Time {
	return a.start
}

func (a Activity) End() time.Time {
	return a.end
}

func (a Activity) CadetStudentFee() uint {
	return a.cadetStudentFee
}

func (a Activity) CadetCadreFee() uint {
	return a.cadetCadreFee
}

func (a Activity) SeniorStudentFee() uint {
	return a.seniorStudentFee
}

func (a Activity) SeniorCadreFee() uint {
	return a.seniorCadreFee
}
