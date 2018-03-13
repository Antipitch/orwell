package orwell

import (
	"fmt"
)

// XOrAnd func
func (*orwell) XAndOr(and interface{}, ors ...interface{}) *xAndOr {
	return &xAndOr{
		and: and,
		ors: ors,
		msg: fmt.Sprintf("Validation error for 'XOrAnd' rule"),
	}
}

type xAndOr struct {
	and interface{}
	ors []interface{}
	msg string
}

// Apply func
func (r *xAndOr) Apply(value interface{}) error {
	if NOE(r.and) || !NOE(value) {
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
