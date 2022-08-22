package config

import (
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func Load() (*Config, error) {
	godotenv.Load()
	configContent, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	configContent = []byte(os.ExpandEnv(string(configContent)))
	config := &Config{}
	if err := yaml.Unmarshal(configContent, config); err != nil {
		panic(err)
	}

	return config, nil
}
