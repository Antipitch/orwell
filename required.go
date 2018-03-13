package orwell

import (
	"fmt"
)

// Required func
func (*orwell) Required() *required {
	return &required{
		msg: fmt.Sprintf("Validation error for 'Required' rule"),
	}
}

// Required struct
type required struct {
	msg string
}

// Apply func
func (r *required) Apply(value interface{}) error {
	if NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
