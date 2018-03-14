package orwell

import (
	"fmt"
)

// XGtOr func
func (*Orwell) XLtOr(xv interface{}, lt int, ors ...interface{}) *xLtOr {
	return &xLtOr{
		xv:  xv,
		lt:  lt,
		ors: ors,
		msg: fmt.Sprintf("Validation error for 'XLtOr' rule"),
	}
}

// xLtOr struct
type xLtOr struct {
	xv  interface{}
	lt  int
	ors []interface{}
	msg string
}

// Apply rule
func (r *xLtOr) Apply(value interface{}) error {
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

	if xvInt < int64(r.lt) {
		for _, or := range r.ors {
			if !NOE(or) {
				return nil
			}
		}
		return fmt.Errorf(r.msg)
	}

	return nil
}
