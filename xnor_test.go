package orwell

import (
	"testing"
)

func TestApplyXNOr(t *testing.T) {
	t.Run("nors not nil, value not nil", func(t *testing.T) {
		r := &xNor{nors: []interface{}{1, 2}, msg: "test"}
		if err := r.Apply("test"); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("nors not nil, value nil", func(t *testing.T) {
		r := &xNor{nors: []interface{}{1, 2}, msg: "test"}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid")
		}
	})
	t.Run("nors empty, value not nil", func(t *testing.T) {
		r := &xNor{nors: []interface{}{}, msg: "test"}
		if err := r.Apply("test"); err == nil {
			t.Error("Expected error")
		}
	})
	t.Run("all nors empty, value not nil", func(t *testing.T) {
		r := &xNor{nors: []interface{}{"", 0, nil}, msg: "test"}
		if err := r.Apply("test"); err == nil {
			t.Error("Expected error")
		}
	})
}
