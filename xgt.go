package orwell

import (
	"fmt"
)

// XGt func
func (*Orwell) XGt(xv interface{}, gt int) *xGt {
	return &xGt{
		xv:  xv,
		gt:  gt,
		msg: fmt.Sprintf("Validation error for 'XGt' rule"),
	}
}

// xGt struct
type xGt struct {
	xv  interface{}
	gt  int
	msg string
}

// Apply rule
func (r *xGt) Apply(value interface{}) error {
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
		return fmt.Errorf(r.msg)
	}

	return nil
}
