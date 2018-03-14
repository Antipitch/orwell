package orwell

import (
	"testing"
)

func TestApplyXGtOr(t *testing.T) {
	t.Run("xv > gt, ors nil, value not nil", func(t *testing.T) {
		r := &xGtOr{xv: 10, gt: 9, ors: nil}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv <= gt, ors nil, value nil", func(t *testing.T) {
		r := &xGtOr{xv: 10, gt: 10, ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv > gt, ors not nil, value nil", func(t *testing.T) {
		r := &xGtOr{xv: 10, gt: 9, ors: []interface{}{"foo"}}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv > gt, ors nil, value nil", func(t *testing.T) {
		r := &xGtOr{xv: 10, gt: 9, ors: nil}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
}
