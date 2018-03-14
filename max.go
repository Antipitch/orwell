package orwell

import (
	"fmt"
)

// Max func
func (*Orwell) Max(arg int) *max {
	return &max{
		max: arg,
		msg: "Validation error for 'Max' rule",
	}
}

type max struct {
	max int
	msg string
}

// Apply func
func (r *max) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(v) {
		return nil
	}

	vi, err := ToInt64(v)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	max := int64(r.max)
	if vi > max {
		return fmt.Errorf(r.msg)
	}

	return nil
}
