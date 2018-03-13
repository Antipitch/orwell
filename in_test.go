package orwell

import (
	"testing"
)

func TestApplyInWithIntegers(t *testing.T) {
	r := in{args: []interface{}{1, 2, 3}, msg: "test"}

	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid int", func(t *testing.T) {
		if err := r.Apply(2); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid int", func(t *testing.T) {
		if err := r.Apply(4); err == nil {
			t.Error("Expected error")
		}
	})
}

func TestApplyInWithStrings(t *testing.T) {
	r := in{args: []interface{}{"1", "2", "3"}, msg: "test"}
	t.Run("Valid string", func(t *testing.T) {
		if err := r.Apply("2"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid string", func(t *testing.T) {
		if err := r.Apply("4"); err == nil {
			t.Error("Expected error")
		}
	})
}
