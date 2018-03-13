package orwell

import (
	"testing"
)

func TestApplyXGt(t *testing.T) {
	t.Run("xv > gt, value not nil", func(t *testing.T) {
		r := &xGt{xv: 10, gt: 9, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv <= gt, value nil", func(t *testing.T) {
		r := &xGt{xv: 10, gt: 10, msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv > gt, value nil", func(t *testing.T) {
		r := &xGt{xv: 10, gt: 9, msg: "test"}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
}
