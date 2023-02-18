package common

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var filepath = "config/config.yaml"

var SiteInfo = GetConfig()

type author struct {
	Name   string `yaml:"name"`
	Avatar string `yaml:"avatar"`
	Url    string `yaml:"url"`
	QQ     string `yaml:"QQ"`
	Email  string `yaml:"email"`
	Github string `yaml:"github"`
}

type site struct {
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	Keywords    string `yaml:"keywords"`
	Description string `yaml:"description"`
	ICP         string `yaml:"ICP"`
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
