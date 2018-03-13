package orwell

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
)

// IsEMail func
func (*orwell) Email(doLookUp bool) *email {
	return &email{
		doLookUp: doLookUp,
		msg:      "Validation error for 'Email' rule",
	}
}

type email struct {
	doLookUp bool
	msg      string
}

// Apply rule
func (r *email) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	vStr, err := ToString(v)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	_, err = mail.ParseAddress(vStr)
	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}
	if r.doLookUp {
		mx, err := net.LookupMX(strings.Split(vStr, "@")[1])
		if err != nil {
			return fmt.Errorf("%s: %s", r.msg, err.Error())
		}
		if len(mx) == 0 {
			return fmt.Errorf("%s: no mx record found", r.msg)
		}
	}

	return nil
}
