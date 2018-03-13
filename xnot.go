package orwell

import (
	"fmt"
)

// XNot func
func (*orwell) XNot(not interface{}) *xNot {
	return &xNot{
		not: not,
		msg: fmt.Sprintf("Validation error for 'XNot' rule"),
	}
}

// xNot struct
type xNot struct {
	not interface{}
	msg string
}

// Apply rule
func (r *xNot) Apply(value interface{}) error {
	if !NOE(r.not) && !NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
