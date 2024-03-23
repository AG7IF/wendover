package org

import (
	"fmt"
)

type Person struct {
	capid     uint
	grade     Grade
	firstName string
	lastName  string
}

func NewPerson(
	capid uint,
	grade Grade,
	firstName string,
	lastName string,
) Person {
	return Person{
		capid:     capid,
		grade:     grade,
		lastName:  lastName,
		firstName: firstName,
	}
}

func (p Person) CAPID() uint {
	return p.capid
}

func (p Person) Grade() Grade {
	return p.grade
}

func (p Person) FirstName() string {
	return p.firstName
}

func (p Person) LastName() string {
	return p.lastName
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s %s", p.grade, p.firstName, p.lastName)
}
