package orwell

import (
	"testing"
)

func TestApplyXNot(t *testing.T) {
	t.Run("Not is not nil, value is nil", func(t *testing.T) {
		r := &xNot{not: "test", msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Not is nil, value is nil", func(t *testing.T) {
		r := &xNot{msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Not is not nil, value is not nil", func(t *testing.T) {
		r := &xNot{not: "test", msg: "test"}
		if err := r.Apply("test"); err == nil {
			t.Error("Expected error")
		}
	})
}
