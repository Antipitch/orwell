package orwell

import (
	"fmt"
)

// XGtAndOr func
func (*Orwell) XGtAndOr(gtField interface{}, gtValue int, and interface{}, ors ...interface{}) *xGtAndOr {
	return &xGtAndOr{
		gtField:  gtField,
		gtValue:  gtValue,
		and: and,
		ors:      ors,
		msg:      fmt.Sprintf("Validation error for 'XGtAndOr' rule"),
	}
}

// xGtAndOr struct
type xGtAndOr struct {
	gtField  interface{}
	gtValue  int
	and interface{}
	ors      []interface{}
	msg      string
}

// Apply rule
func (r *xGtAndOr) Apply(value interface{}) error {
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

	if gtFieldInt64 > int64(r.gtValue) && !NOE(r.and) {
		for _, or := range r.ors {
			if !NOE(or) {
				return nil
			}
		}
		return fmt.Errorf(r.msg)
	}

	return nil
}
