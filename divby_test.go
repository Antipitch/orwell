package orwell

import (
	"testing"
)

func TestApplyDivBy(t *testing.T) {
	r := &divBy{div: 10, msg: "test"}
	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid int", func(t *testing.T) {
		if err := r.Apply(100); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid int", func(t *testing.T) {
		if err := r.Apply(25); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Pointer", func(t *testing.T) {
		i := 100
		if err := r.Apply(&i); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Pointer", func(t *testing.T) {
		i := 25
		if err := r.Apply(&i); err == nil {
			t.Error("Expected error")
		}
	})
}
