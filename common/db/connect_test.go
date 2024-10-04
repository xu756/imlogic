package db

import (
	"testing"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	if m == nil {
		t.Error("NewModel error")
	}
}
