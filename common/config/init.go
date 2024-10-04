package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var runData Config

func RunData() Config {
	// 判断是否为空
	if runData == (Config{}) {
		configFileBytes, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
		if err != nil {
			panic(err)
		}
		d := Config{}
		err = yaml.Unmarshal(configFileBytes, &d)
		if err != nil {
			panic(err)
		}
		runData = d
	}
	return runData

}

func InitDir() {
	err := CreateDir(RunData().UploadPath)
	if err != nil {
		return
	}
}
