package orwell

import (
	"fmt"
)

// XLt func
func (*orwell) XLt(v interface{}, lt int) *XLt {
	return &XLt{
		v:   v,
		lt:  lt,
		msg: "Validation error for 'XGt' rule",
	}
}

// XLt struct
type XLt struct {
	v   interface{}
	lt  int
	msg string
}

// Apply rule
func (r *XLt) Apply(value interface{}) error {
	v, err := ToInt64(r.v)

	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if v < int64(r.lt) && NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
