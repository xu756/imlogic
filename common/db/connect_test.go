package db

import (
	"imlogic/common/config"
	"testing"
)

func TestNewModel(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	if m == nil {
		t.Error("NewModel error")
	}
}
