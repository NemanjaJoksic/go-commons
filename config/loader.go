package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func LoadConfiguration[T any](config *T, filePath string) {
	loadConfigurationFromYamlFile(config, filePath)
	loadConfigurationFromEnvVars(config)
}

func loadConfigurationFromYamlFile[T any](config *T, filePath string) {
	fmt.Println("Read config ", filePath)

	file, err1 := os.Open(filePath)

	if err1 != nil {
		fmt.Println("err1: ", err1)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err2 := decoder.Decode(config)
	if err2 != nil {
		fmt.Println("err2: ", err2)
	}
}

func loadConfigurationFromEnvVars[T any](config *T) {
	err := envconfig.Process("", config)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
