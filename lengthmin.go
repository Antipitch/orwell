package orwell

import (
	"fmt"
	"unicode/utf8"
)

// LengthMin func
func (*Orwell) LengthMin(min int) *lengthMin {
	return &lengthMin{
		min: min,
		msg: "Validation error for 'LengthMin' rule",
	}
}

type lengthMin struct {
	min int
	msg string
}

// Apply func
func (r *lengthMin) Apply(value interface{}) error {
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

	if l < r.min {
		return fmt.Errorf(r.msg)
	}

	return nil
}
