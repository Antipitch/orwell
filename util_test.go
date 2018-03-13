package orwell

import (
	"testing"
)

func TestNoe(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		if NOE("test") {
			t.Error("Expected false")
		}
	})
	t.Run("StringEmpty", func(t *testing.T) {
		var test string
		if !NOE(test) {
			t.Error("Expected true")
		}
	})
	t.Run("StringDefault", func(t *testing.T) {
		if !NOE("") {
			t.Error("Expected true")
		}
	})
	t.Run("Int", func(t *testing.T) {
		if NOE(1) {
			t.Error("Expected false")
		}
	})
	t.Run("IntEmpty", func(t *testing.T) {
		var test int
		if !NOE(test) {
			t.Error("Expected true")
		}
	})
	t.Run("IntDefault", func(t *testing.T) {
		if !NOE(0) {
			t.Error("Expected true")
		}
	})
	t.Run("Map", func(t *testing.T) {
		m := make(map[int]int)
		m[0] = 1
		m[1] = 2
		if NOE(m) {
			t.Error("Expected false")
		}
	})
	t.Run("MapEmpty", func(t *testing.T) {
		m := make(map[int]int)
		if !NOE(m) {
			t.Error("Expected true")
		}
	})
}
