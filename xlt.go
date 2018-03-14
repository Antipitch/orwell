package orwell

import (
	"fmt"
)

// XLt func
func (*Orwell) XLt(v interface{}, lt int) *xLt {
	return &xLt{
		xv:  v,
		lt:  lt,
		msg: "Validation error for 'xGt' rule",
	}
}

// xLt struct
type xLt struct {
	xv  interface{}
	lt  int
	msg string
}

// Apply rule
func (r *xLt) Apply(value interface{}) error {
	xv, isNil := IsNil(r.xv)
	if isNil || IsEmpty(xv) {
		return nil
	}

	xvInt, err := ToInt64(xv)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if xvInt < int64(r.lt) && NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
