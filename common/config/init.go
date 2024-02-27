package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var RunData Config

func Init(path string) {
	configFileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	d := Config{}
	err = yaml.Unmarshal(configFileBytes, &d)
	if err != nil {
		panic(err)
	}
	RunData = d
}
