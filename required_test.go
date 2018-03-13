package orwell

import (
	"testing"
)

func TestApplyRequired(t *testing.T) {
	r := &required{msg: "test"}
	t.Run("Nil value", func(t *testing.T) {
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid int", func(t *testing.T) {
		if err := r.Apply(1); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid int", func(t *testing.T) {
		if err := r.Apply(0); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid string", func(t *testing.T) {
		if err := r.Apply("1"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid string", func(t *testing.T) {
		if err := r.Apply(""); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Pointer", func(t *testing.T) {
		i := 1
		if err := r.Apply(&i); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Pointer", func(t *testing.T) {
		var i *int
		if err := r.Apply(i); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Map", func(t *testing.T) {
		m := make(map[string]int)
		m["test"] = 1
		if err := r.Apply(m); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Map", func(t *testing.T) {
		m := make(map[string]int)
		if err := r.Apply(m); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("Valid Struct", func(t *testing.T) {
		type TestStruct struct {
			test int
		}
		if err := r.Apply(TestStruct{test: 1}); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Invalid Struct", func(t *testing.T) {
		type TestStruct struct {
			test int
		}
		if err := r.Apply(TestStruct{}); err == nil {
			t.Error("Expected error")
		}
	})

}
