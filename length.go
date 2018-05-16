package orwell

import (
	"fmt"
	"unicode/utf8"
)

// Length func
func (*Orwell) Length(min, max int) *length {
	return &length{
		min: min,
		max: max,
		msg: "Validation error for 'length' rule",
	}
}

type length struct {
	min, max int
	msg      string
}

// Apply func
func (r *length) Apply(value interface{}) error {
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

	if l < r.min || l > r.max {
		return fmt.Errorf(r.msg)
	}

	return nil
}
