package orwell

import (
	"fmt"
)

// XAnd func
func (*orwell) XAnd(and interface{}) *xAnd {
	return &xAnd{
		and: and,
		msg: fmt.Sprintf("Validation error for 'XAnd' rule"),
	}
}

// xAnd struct
type xAnd struct {
	and interface{}
	msg string
}

// Apply func
func (r *xAnd) Apply(value interface{}) error {
	if !NOE(r.and) && !NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
