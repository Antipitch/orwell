package orwell

import (
	"fmt"
)

// XNot func
func (*orwell) XNot(not interface{}) *XNot {
	return &XNot{
		not: not,
		msg: fmt.Sprintf("Validation error for 'XNot' rule"),
	}
}

// XNot struct
type XNot struct {
	not interface{}
	msg string
}

// Apply rule
func (r *XNot) Apply(value interface{}) error {
	if !NOE(r.not) && !NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
