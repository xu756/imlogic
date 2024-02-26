package db

import (
	"github.com/xu756/imlogic/common/config"
	"testing"
)

func TestNewModel(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	if m == nil {
		t.Error("NewModel error")
	}
}
