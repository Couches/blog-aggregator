package blogaggregator

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfiguration struct {
  Server struct {
    Address string `yaml:"address"`
  } `yaml:"server"`
  Database struct {
    ConnectionString string `yaml:"connection_string"`
  } `yaml:"database"`
}

func LoadAppConfiguration(path string) (AppConfiguration, error) {
  config := AppConfiguration{}
  file, err := os.ReadFile(path)
  if err != nil {
    return config, err
  }

  err = yaml.Unmarshal(file, &config)
  if err != nil {
    return config, err
  }

  return config, nil
}
