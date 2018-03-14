package orwell

import (
	"fmt"
)

// XGtOr func
func (*Orwell) XGtOr(xv interface{}, gt int, ors ...interface{}) *xGtOr {
	return &xGtOr{
		xv:  xv,
		gt:  gt,
		ors: ors,
		msg: fmt.Sprintf("Validation error for 'XGtOr' rule"),
	}
}

// xGtOr struct
type xGtOr struct {
	xv  interface{}
	gt  int
	ors []interface{}
	msg string
}

// Apply rule
func (r *xGtOr) Apply(value interface{}) error {
	if !NOE(value) {
		return nil
	}

	xv, isNil := IsNil(r.xv)
	if isNil || IsEmpty(xv) {
		return nil
	}

	xvInt, err := ToInt64(xv)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if xvInt > int64(r.gt) {
		for _, or := range r.ors {
			if !NOE(or) {
				return nil
			}
		}
		return fmt.Errorf(r.msg)
	}

	return nil
}
