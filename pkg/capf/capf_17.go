package capf

import (
	"time"

	"github.com/ag7if/wendover/pkg/personnel"
)

type Capf17 struct {
	// Items 1-7
	ActivityTitle            string
	ActivityLocation         string
	ActivityDates            string
	PreviouslyAttended       bool
	DateOfPreviousAttendance time.Time
	Name                     string
	Grade                    personnel.Grade
	CAPID                    uint

	// Items 7-15
	Address1             string
	Address2             string
	CityStateAndZip      string
	WorkPhone            string
	HomePhone            string
	Email                string
	HomeUnit             personnel.HomeUnit
	Level1CompletionDate time.Time
	CapJoinDate          time.Time
	CapDutyAssignment    string
	CapAeroRating        string

	// Item 16
	SpecialtyTrackA string
	RatingA         string
	SpecialtyTrackB string
	RatingB         string
	SpecialtyTrackC string
	RatingC         string
	SpecialtyTrackD string
	RatingD         string

	// Item 17
	TrainingActivityA string
	TrainingActivityB string
	TrainingActivityC string
	TrainingActivityD string
	TrainingActivityE string

	// Item 18
	ProfessionalDevelopmentAwardA string
	ProfessionalDevelopmentAwardB string
	ProfessionalDevelopmentAwardC string
	ProfessionalDevelopmentAwardD string

	// Item 19
	HighSchoolGraduationYear   uint
	CollegeYearsCompleted      uint
	PostGraduateYearsCompleted uint

	// Items 20-23
	CivilianOccupation          string
	EmergencyMedicalInformation string
	PersonalCapGoalsOutline     string
	Remarks                     string

	// Items 24-26
	UnitCCApproved  bool
	UnitCCRemarks   string
	UnitCCSignature string

	WingCCApproved  bool
	WingCCRemarks   string
	WingCCSignature string

	RegionCCApproved  bool
	RegionCCRemarks   string
	RegionCCSignature string

	// Item 27
	AdditionalRemarks string
}
