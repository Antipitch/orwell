package orwell

import (
	"testing"
)

func TestApplyXGtIsOr(t *testing.T) {
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtIsOr{gtField: 10, gtValue: 9, isField: "test", isValue: "test", ors: nil}
		if err := r.Apply("not nil"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtIsOr{gtField: 10, gtValue: 9, isField: "test", isValue: "test", ors: nil}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtIsOr{gtField: 10, gtValue: 10, isField: 0, isValue: 0, ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because gtField is not greater than igtValue")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtIsOr{gtField: 10, gtValue: 9, isField: 0, isValue: 1, ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because isField and isValue are not equal")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtIsOr{gtField: 10, gtValue: 9, isField: 1, isValue: 1, ors: []interface{}{"test"}}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because ors is not empty")
		}
	})
}
