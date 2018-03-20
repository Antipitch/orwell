package orwell

import (
	"fmt"
	"regexp"
)

// Match func
func (*Orwell) Match(arg string) *match {
	return &match{
		regexp: regexp.MustCompile(arg),
		msg:    "Validation error for 'Match' rule",
	}
}

type match struct {
	regexp *regexp.Regexp
	msg    string
}

// Apply func
func (r *match) Apply(value interface{}) error {
	v, isNil := IsNil(value)
	if isNil || IsEmpty(v) {
		return nil
	}

	vs, err := ToString(v)

	if err != nil {
		return fmt.Errorf("%s: %s", r.msg, err.Error())
	}

	matched := r.regexp.MatchString(vs)

	if matched != true {
		return fmt.Errorf(r.msg)
	}
	return nil
}
