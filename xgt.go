package orwell

import (
	"fmt"
)

// XGt func
func (*orwell) XGt(v interface{}, gt int) *XGt {
	return &XGt{
		v:   v,
		gt:  gt,
		msg: fmt.Sprintf("Validation error for 'XGt' rule"),
	}
}

// XGt struct
type XGt struct {
	v   interface{}
	gt  int
	msg string
}

// Apply rule
func (r *XGt) Apply(value interface{}) error {
	v, err := ToInt64(r.v)

	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if v > int64(r.gt) && NOE(value) {
		return fmt.Errorf(r.msg)
	}

	return nil
}
