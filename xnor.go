package orwell

import (
	"fmt"
)

// XNor func
func (*orwell) XNor(nor interface{}) *XNor {
	return &XNor{
		nor: nor,
		msg: fmt.Sprintf("Validation error for 'XNor' rule"),
	}
}

// XNor struct
type XNor struct {
	nor interface{}
	msg string
}

// Apply func
func (r *XNor) Apply(value interface{}) error {
	if NOE(r.nor) && !NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
