package orwell

import (
	"fmt"
)

// In func
func (*orwell) In(args ...interface{}) *in {
	return &in{
		args: args,
		msg:  fmt.Sprintf("Validation error for 'In' rule"),
	}
}

// In struct
type in struct {
	args []interface{}
	msg  string
}

// Apply rule
func (r *in) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(v) {
		return nil
	}

	switch v.(type) {
	case string:
		vs, _ := ToString(v)
		for _, arg := range r.args {
			if as, err := ToString(arg); err == nil {
				if vs == as {
					return nil
				}
			}
		}
	case int, int8, int16, int32, int64:
		vi, _ := ToInt64(v)
		for _, arg := range r.args {
			if ai, err := ToInt64(arg); err == nil {
				if vi == ai {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("%s: value and rule argument types are not comparable", r.msg)
}
