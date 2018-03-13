package orwell

import (
	"fmt"
)

// XNor func
func (*orwell) XNor(nors ...interface{}) *xNor {
	return &xNor{
		nors: nors,
		msg:  fmt.Sprintf("Validation error for 'XNor' rule"),
	}
}

// xNor struct
type xNor struct {
	nors []interface{}
	msg  string
}

// Apply func
func (r *xNor) Apply(value interface{}) error {
	if NOE(value) {
		return nil
	}

	if len(r.nors) == 0 {
		return fmt.Errorf(r.msg)
	}

	for _, nor := range r.nors {
		if !NOE(nor) {
			return nil
		}
	}

	return fmt.Errorf(r.msg)
}
