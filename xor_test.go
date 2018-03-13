package orwell

import (
	"testing"
)

func TestApplyXOr(t *testing.T) {
	t.Run("Ors not nil, value not nil", func(t *testing.T) {
		r := &xOr{ors: []interface{}{1, 2}, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})

	t.Run("All ors empty, value not nil", func(t *testing.T) {
		r := &xOr{ors: []interface{}{"", 0, nil}, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected error")
		}
	})
	t.Run("Ors not nil, value nil", func(t *testing.T) {
		r := &xOr{ors: []interface{}{"", 2}, msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("Ors empty, value nil", func(t *testing.T) {
		r := &xOr{ors: []interface{}{}, msg: "test"}
		if err := r.Apply(""); err == nil {
			t.Error("Expected error")
		}
	})
}
