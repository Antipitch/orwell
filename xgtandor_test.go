package orwell

import (
	"testing"
)

func TestApplyXGtAndOr(t *testing.T) {
	t.Run("gtField > gtValue, andValue not empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 9, and: "not nil", ors: nil}
		if err := r.Apply("not nil"); err != nil {
			t.Error("Expected valid because validated value is not nil ")
		}
	})
	t.Run("gtField > gtValue, andValue not empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 9, and: "not nil", ors: nil}
		if err := r.Apply(nil); err == nil {
			t.Error("Expected error because validated value is nil")
		}
	})
	t.Run("gtField > gtValue, andValue is empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 9, and: nil, ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because andValue is nil")
		}
	})
	t.Run("gtField > gtValue, andValue is empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 9, and: "", ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because andValue is empty")
		}
	})
	t.Run("gtField = gtValue, andValue not empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 10, and: "not nil", ors: nil}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because gtField is not greater than gtValue")
		}
	})

	t.Run("gtField > gtValue, andValue not empty, ors nil, value not nil", func(t *testing.T) {
		r := &xGtAndOr{gtField: 10, gtValue: 9, and: "not nil", ors: []interface{}{"test"}}
		if err := r.Apply(nil); err != nil {
			t.Error("Expected valid because ors is not empty")
		}
	})
}
