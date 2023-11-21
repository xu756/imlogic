package config

import (
	"log"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := CreateDir("./test/11/11")
	if err != nil {
		log.Print(err)
		return
	}
}
