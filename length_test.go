package orwell

import (
	"testing"
)

func TestApplyLength(t *testing.T) {
	r := &length{min: 2, max: 4, msg: "test"}
	t.Run("String invalid too short", func(t *testing.T) {
		if err := r.Apply("ö"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("String valid equal to min", func(t *testing.T) {
		if err := r.Apply("öä"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String valid between min and max", func(t *testing.T) {
		if err := r.Apply("ööö"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String valid equal to max", func(t *testing.T) {
		if err := r.Apply("öööö"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("String invalid too long", func(t *testing.T) {
		if err := r.Apply("ööööö"); err == nil {
			t.Error("Expected error")
		}
	})
}
