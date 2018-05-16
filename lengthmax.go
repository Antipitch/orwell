package orwell

import (
	"fmt"
	"unicode/utf8"
)

// LengthMax func
func (*Orwell) LengthMax(max int) *lengthMax {
	return &lengthMax{
		max: max,
		msg: "Validation error for 'LengthMax' rule",
	}
}

type lengthMax struct {
	max int
	msg string
}

// Apply func
func (r *lengthMax) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(v) {
		return nil
	}

	var (
		l   int
		err error
	)
	if str, ok := value.(string); ok {
		l = utf8.RuneCountInString(str)
	} else if l, err = LengthOf(value); err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if l > r.max {
		return fmt.Errorf(r.msg)
	}

	return nil
}
