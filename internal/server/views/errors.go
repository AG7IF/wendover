package views

type ErrInvalidValue struct {
	FieldName   string
	ValidValues []string
}

func (err ErrInvalidValue) Error() string {
	return "invalid value set on field"
}
