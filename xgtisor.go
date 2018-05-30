package orwell

import (
	"fmt"
)

// XGtIsOr func
func (*Orwell) XGtIsOr(gtField interface{}, gtValue int, isField interface{}, isValue interface{}, ors ...interface{}) *xGtIsOr {
	return &xGtIsOr{
		gtField: gtField,
		gtValue: gtValue,
		isField: isField,
		isValue: isValue,
		ors:     ors,
		msg:     fmt.Sprintf("Validation error for 'XGtOr' rule"),
	}
}

// xGtOr struct
type xGtIsOr struct {
	gtField interface{}
	gtValue int
	isField interface{}
	isValue interface{}
	ors     []interface{}
	msg     string
}

// Apply rule
func (r *xGtIsOr) Apply(value interface{}) error {
	if !NOE(value) {
		return nil
	}

	gtField, isNil := IsNil(r.gtField)
	if isNil || IsEmpty(gtField) {
		return nil
	}

	gtFieldInt64, err := ToInt64(gtField)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	if gtFieldInt64 > int64(r.gtValue) && DeepEqual(r.isField, r.isValue) {
		for _, or := range r.ors {
			if !NOE(or) {
				return nil
			}
		}
		return fmt.Errorf(r.msg)
	}

	return nil
}
