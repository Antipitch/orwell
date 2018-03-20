package orwell

import (
	"regexp"
	"testing"
)

func TestApplyMatch(t *testing.T) {
	r := &match{regexp: regexp.MustCompile("^te[s]t"), msg: "Regexp test"}
	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Valid string", func(t *testing.T) {
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid string", func(t *testing.T) {
		if err := r.Apply(" tet"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Pointer", func(t *testing.T) {
		s := "test2"
		if err := r.Apply(&s); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Pointer", func(t *testing.T) {
		s := "2test"
		if err := r.Apply(&s); err == nil {
			t.Error("Expected error")
		}
	})
}
