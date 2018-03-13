package orwell

import (
	"fmt"
)

// XGt func
func (*orwell) XGt(v interface{}, gt int) *xGt {
	return &xGt{
		xv:  v,
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
	xv, isNil := IsNil(r.xv)
	if isNil || IsEmpty(xv) {
		return nil
	}

	xvInt, err := ToInt64(xv)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if xvInt > int64(r.gt) && NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
