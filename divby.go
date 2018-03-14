package orwell

import (
	"fmt"
)

// DivBy func
func (*Orwell) DivBy(div int) *divBy {
	return &divBy{
		div: div,
		msg: "Validation error for 'DivBy' rule",
	}
}

type divBy struct {
	div int
	msg string
}

// Apply func
func (r *divBy) Apply(value interface{}) error {
	v, isNil := IsNil(value)

	if isNil || IsEmpty(value) {
		return nil
	}

	dd, err := ToInt64(v)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	div := int64(r.div)
	if dd%div != 0 {
		return fmt.Errorf(r.msg)
	}

	return nil
}
