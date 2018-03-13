package orwell

import (
	"fmt"
)

// XOr func
func (*orwell) XOr(ors ...interface{}) *xOr {
	return &xOr{
		ors: ors,
		msg: "Validation error for 'XOr' rule",
	}
}

// xOr struct
type xOr struct {
	ors []interface{}
	msg string
}

// Apply rule
func (r *xOr) Apply(value interface{}) error {
	if !NOE(value) {
		return nil
	}

	if len(r.ors) == 0 {
		return fmt.Errorf(r.msg)
	}

	for _, or := range r.ors {
		if !NOE(or) {
			return nil
		}
	}

	return fmt.Errorf(r.msg)
}
