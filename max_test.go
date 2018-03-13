package orwell

import (
	"testing"
)

func TestApplyMax(t *testing.T) {
	r := &max{max: 10, msg: "test"}
	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid int", func(t *testing.T) {
		if err := r.Apply(9); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid int equal", func(t *testing.T) {
		if err := r.Apply(10); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid int", func(t *testing.T) {
		if err := r.Apply(11); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Pointer", func(t *testing.T) {
		i := 9
		if err := r.Apply(&i); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Pointer", func(t *testing.T) {
		i := 11
		if err := r.Apply(&i); err == nil {
			t.Error("Expected error")
		}
	})
}
