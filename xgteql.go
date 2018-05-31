package orwell

import (
	"fmt"
)

// XGtEql func
func (*Orwell) XGtEql(gtField interface{}, gtValue int, eqlField interface{}, eqlValue interface{}) *xGtEql {
	return &xGtEql{
		gtField:  gtField,
		gtValue:  gtValue,
		eqlField: eqlField,
		eqlValue: eqlValue,
		msg:      fmt.Sprintf("Validation error for 'XGtEql' rule"),
	}
}

// xGtEql struct
type xGtEql struct {
	gtField  interface{}
	gtValue  int
	eqlField interface{}
	eqlValue interface{}
	msg      string
}

// Apply rule
func (r *xGtEql) Apply(value interface{}) error {
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
		return fmt.Errorf(r.msg)
	}

	return nil
}
