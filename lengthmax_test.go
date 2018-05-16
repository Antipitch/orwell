package orwell

import (
	"testing"
)

func TestApplyLengthMax(t *testing.T) {
	r := &lengthMax{max: 2, msg: "test"}
	t.Run("String valid length less than max", func(t *testing.T) {
		if err := r.Apply("ö"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String valid length equal to min", func(t *testing.T) {
		if err := r.Apply("öä"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String invalid length greater than max", func(t *testing.T) {
		if err := r.Apply("öäü"); err == nil {
			t.Error("Expected error")
		}
	})
}
