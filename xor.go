package orwell

import (
	"fmt"
)

// XOr func
func (*orwell) XOr(ors ...interface{}) *XOr {
	return &XOr{
		ors: ors,
		msg: "Validation error for 'XOr' rule",
	}
}

// XOr struct
type XOr struct {
	ors []interface{}
	msg string
}

// Apply rule
func (r *XOr) Apply(value interface{}) error {
	if !NOE(value) {
		return nil
	}

	if len(r.ors) == 0 {
		return nil
	}

	for _, or := range r.ors {
		if !NOE(or) {
			return nil
		}
	}

	return fmt.Errorf(r.msg)
}
