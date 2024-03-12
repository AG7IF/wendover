package views

import (
	"time"

	"github.com/google/uuid"

	"github.com/ag7if/wendover/pkg/org"
)

type ActivityView struct {
	ID               uuid.UUID `json:"id"`
	Key              string    `json:"key"`
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
	CadetStudentFee  uint      `json:"cadet_student_fee"`
	CadetCadreFee    uint      `json:"cadet_cadre_fee"`
	SeniorStudentFee uint      `json:"senior_student_fee"`
	SeniorCadreFee   uint      `json:"senior_cadre_fee"`
}

func NewActivityView(activity org.Activity) ActivityView {
	return ActivityView{
		ID:               activity.ID(),
		Key:              activity.Key(),
		Name:             activity.Name(),
		Location:         activity.Location(),
		Start:            activity.Start(),
		End:              activity.End(),
		CadetStudentFee:  activity.CadetStudentFee(),
		CadetCadreFee:    activity.CadetCadreFee(),
		SeniorStudentFee: activity.SeniorStudentFee(),
		SeniorCadreFee:   activity.SeniorCadreFee(),
	}
}

func (av ActivityView) ToDomainObject() org.Activity {
	a := org.NewActivity(
		av.ID,
		av.Key,
		av.Name,
		av.Location,
		av.Start,
		av.End,
		av.CadetStudentFee,
		av.CadetCadreFee,
		av.SeniorStudentFee,
		av.SeniorCadreFee,
	)

	return a
}
