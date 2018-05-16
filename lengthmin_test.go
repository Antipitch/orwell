package orwell

import (
	"testing"
)

func TestApplyLengthMin(t *testing.T) {
	r := &lengthMin{min: 2, msg: "test"}
	t.Run("String invalid length less than min", func(t *testing.T) {
		if err := r.Apply("ö"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("String valid length equal to min", func(t *testing.T) {
		if err := r.Apply("öä"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String valid length greater than min", func(t *testing.T) {
		if err := r.Apply("öä"); err != nil {
			t.Error("Expected valid")
		}
	})
}
