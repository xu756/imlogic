package config

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	Init("../../configs/dev.yaml")
	log.Print(RunData)
}
