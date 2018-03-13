package orwell

import (
	"testing"
)

func TestApplyXLt(t *testing.T) {
	t.Run("xv < gt, value not nil", func(t *testing.T) {
		r := &xLt{v: 9, lt: 10, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv >= gt, value nil", func(t *testing.T) {
		r := &xLt{v: 10, lt: 10, msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("xv < gt, value nil", func(t *testing.T) {
		r := &xLt{v: 9, lt: 10, msg: "test"}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error")
		}
	})
}
