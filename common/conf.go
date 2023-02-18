package common

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var filepath = "conf/config.yaml"

type author struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	QQ   string `yaml:"QQ"`
}

type site struct {
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Config struct {
	Author author `yaml:"author"`
	Site   site   `yaml:"site"`
}

func GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		Log("app").Error().Msg(err.Error())
	}
	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		Log("app").Error().Msg(err.Error())
	}
	return &config
}
