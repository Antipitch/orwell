package orwell

import (
	"fmt"
)

// Max func
func (*Orwell) Min(arg int) *min {
	return &min{
		min: arg,
		msg: "Validation error for 'Min' rule",
	}
}

type min struct {
	min int
	msg string
}

// Apply func
func (r *min) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(v) {
		return nil
	}

	vi, err := ToInt64(v)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	min := int64(r.min)
	if vi < min {
		return fmt.Errorf(r.msg)
	}

	return nil
}
