package orwell

import (
	"testing"
)

func TestApplyXLtOr(t *testing.T) {
	t.Run("xv < glt, ors nil, value not nil", func(t *testing.T) {
		r := &xLtOr{xv: 9, lt: 10, ors: nil}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv >= lt, ors nil, value nil", func(t *testing.T) {
		r := &xLtOr{xv: 10, lt: 10, ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv < lt, ors not nil, value nil", func(t *testing.T) {
		r := &xLtOr{xv: 10, lt: 9, ors: []interface{}{"foo"}}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv < lt, ors nil, value nil", func(t *testing.T) {
		r := &xLtOr{xv: 9, lt: 10, ors: nil}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
}
