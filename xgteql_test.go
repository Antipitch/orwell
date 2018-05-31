package orwell

import (
	"testing"
)

func TestApplyXGtEql(t *testing.T) {
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtEql{gtField: 10, gtValue: 9, eqlField: "test", eqlValue: "test"}
		if err := r.Apply("not nil"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtEql{gtField: 10, gtValue: 9, eqlField: "test", eqlValue: "test"}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtEql{gtField: 10, gtValue: 10, eqlField: 0, eqlValue: 0}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because gtField is not greater than igtValue")
		}
	})
	t.Run("gtField > gtValue, isField == isValue, ors nil, value not nil", func(t *testing.T) {
		r := &xGtEql{gtField: 10, gtValue: 9, eqlField: 0, eqlValue: 1}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because isField and isValue are not equal")
		}
	})
}
