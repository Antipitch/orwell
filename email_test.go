package orwell

import (
	"testing"
)

func TestApplyEmail(t *testing.T) {
	r := &email{doLookUp: false, msg: "test"}
	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid string", func(t *testing.T) {
		if err := r.Apply("test@example.com"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid string no mx", func(t *testing.T) {
		if err := r.Apply("testexample@"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Invalid string no name", func(t *testing.T) {
		if err := r.Apply("@example.com"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Invalid string no @", func(t *testing.T) {
		if err := r.Apply("test.example.com"); err == nil {
			t.Error("Expected error")
		}
	})
}
