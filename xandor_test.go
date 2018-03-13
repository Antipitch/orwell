package orwell

import (
	"testing"
)

func TestApplyXAndOr(t *testing.T) {
	t.Run("And not nil, ors not nil, value not nil", func(t *testing.T) {
		r := &xAndOr{and: "not nil", ors: []interface{}{1, 2}, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("And not nil, ors not nil, value nil", func(t *testing.T) {
		r := &xAndOr{and: "not nil", ors: []interface{}{1, 2}, msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("And not nil, ors nil, value not nil", func(t *testing.T) {
		r := &xAndOr{and: "not nil", msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("And nil, ors nil, value nil", func(t *testing.T) {
		r := &xAndOr{msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("And not nil, ors nil, value nil", func(t *testing.T) {
		r := &xAndOr{and: "not nil", msg: "test"}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
}
