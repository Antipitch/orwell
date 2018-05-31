package orwell

import (
	"fmt"
)

// XGtEqlOr func
func (*Orwell) XGtEqlOr(gtField interface{}, gtValue int, eqlField interface{}, eqlValue interface{}, ors ...interface{}) *xGtEqlOr {
	return &xGtEqlOr{
		gtField:  gtField,
		gtValue:  gtValue,
		eqlField: eqlField,
		eqlValue: eqlValue,
		ors:      ors,
		msg:      fmt.Sprintf("Validation error for 'XGtEqlOr' rule"),
	}
}

// xGtEqlOr struct
type xGtEqlOr struct {
	gtField  interface{}
	gtValue  int
	eqlField interface{}
	eqlValue interface{}
	ors      []interface{}
	msg      string
}

// Apply rule
func (r *xGtEqlOr) Apply(value interface{}) error {
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

	if gtFieldInt64 > int64(r.gtValue) && DeepEqual(r.eqlField, r.eqlValue) {
		for _, or := range r.ors {
			if !NOE(or) {
				return nil
			}
		}
		return fmt.Errorf(r.msg)
	}

	return nil
}
