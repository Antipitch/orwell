package orwell

import (
	"fmt"
)

// XLt func
func (*orwell) XLt(v interface{}, lt int) *xLt {
	return &xLt{
		v:   v,
		lt:  lt,
		msg: "Validation error for 'xGt' rule",
	}
}

// xLt struct
type xLt struct {
	v   interface{}
	lt  int
	msg string
}

// Apply rule
func (r *xLt) Apply(value interface{}) error {
	v, err := ToInt64(r.v)

	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if v < int64(r.lt) && NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
