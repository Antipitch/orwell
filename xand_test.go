package orwell

import (
	"testing"
)

func TestApplyXAnd(t *testing.T) {
	r := &xAnd{and: "not nil", msg: "test"}
	t.Run("And not nil, value not nil", func(t *testing.T) {
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("And not nil, value nil", func(t *testing.T) {
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
	r = &xAnd{msg: "test"}
	t.Run("And nil, value nil", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
}
