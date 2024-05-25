package personnel

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	id                      uuid.UUID
	capid                   uint
	lastName                string
	firstName               string
	middleName              string
	memberType              MemberType
	grade                   Grade
	gender                  Gender
	ageAtEventStart         uint
	ageAtEventEnd           uint
	shirtSize               ShirtSize
	membershipExpires       time.Time
	eServicesEmail          string
	unitApprovalDate        *time.Time
	cadetParentPrimaryPhone *string
	cadetParentPrimaryEmail *string
	unitCCName              string
	unitCCEmail             string
	wingCCName              string
	wingCCEmail             string
	cpptExpires             time.Time
	icutDate                time.Time
	capDlExpires            time.Time
}

func NewParticipant(
	id uuid.UUID,
	capid uint,
	lastName string,
	firstName string,
	middleName string,
	memberType MemberType,
	grade Grade,
	gender Gender,
	ageAtEventStart uint,
	ageAtEventEnd uint,
	shirtSize ShirtSize,
	membershipExpires time.Time,
	eServicesEmail string,
	unitApprovalDate *time.Time,
	cadetParentPrimaryPhone *string,
	cadetParentPrimaryEmail *string,
	unitCCName string,
	unitCCEmail string,
	wingCCName string,
	wingCCEmail string,
	cpptExpires time.Time,
	icutDate time.Time,
	capDlExpires time.Time,
) Participant {
	return Participant{
		id:                      id,
		capid:                   capid,
		lastName:                lastName,
		firstName:               firstName,
		middleName:              middleName,
		memberType:              memberType,
		grade:                   grade,
		gender:                  gender,
		ageAtEventStart:         ageAtEventStart,
		ageAtEventEnd:           ageAtEventEnd,
		shirtSize:               shirtSize,
		membershipExpires:       membershipExpires,
		eServicesEmail:          eServicesEmail,
		unitApprovalDate:        unitApprovalDate,
		cadetParentPrimaryPhone: cadetParentPrimaryPhone,
		cadetParentPrimaryEmail: cadetParentPrimaryEmail,
		unitCCName:              unitCCName,
		unitCCEmail:             unitCCEmail,
		wingCCName:              wingCCName,
		wingCCEmail:             wingCCEmail,
		cpptExpires:             cpptExpires,
		icutDate:                icutDate,
		capDlExpires:            capDlExpires,
	}
}

func (p Participant) Id() uuid.UUID {
	return p.id
}

func (p Participant) CAPID() uint {
	return p.capid
}

func (p Participant) LastName() string {
	return p.lastName
}

func (p Participant) FirstName() string {
	return p.firstName
}

func (p Participant) MiddleName() string {
	return p.middleName
}

func (p Participant) MemberType() MemberType {
	return p.memberType
}

func (p Participant) Grade() Grade {
	return p.grade
}

func (p Participant) Gender() Gender {
	return p.gender
}

func (p Participant) AgeAtEventStart() uint {
	return p.ageAtEventStart
}

func (p Participant) AgeAtEventEnd() uint {
	return p.ageAtEventEnd
}

func (p Participant) ShirtSize() ShirtSize {
	return p.shirtSize
}

func (p Participant) MembershipExpires() time.Time {
	return p.membershipExpires
}

func (p Participant) EServicesEmail() string {
	return p.eServicesEmail
}

func (p Participant) UnitApprovalDate() *time.Time {
	return p.unitApprovalDate
}

func (p Participant) CadetParentPrimaryPhone() *string {
	return p.cadetParentPrimaryPhone
}

func (p Participant) CadetParentPrimaryEmail() *string {
	return p.cadetParentPrimaryEmail
}

func (p Participant) UnitCCName() string {
	return p.unitCCName
}

func (p Participant) UnitCCEmail() string {
	return p.unitCCEmail
}

func (p Participant) WingCCName() string {
	return p.wingCCName
}

func (p Participant) WingCCEmail() string {
	return p.wingCCEmail
}

func (p Participant) CpptExpires() time.Time {
	return p.cpptExpires
}

func (p Participant) IcutDate() time.Time {
	return p.icutDate
}

func (p Participant) CapDlExpires() time.Time {
	return p.capDlExpires
}
