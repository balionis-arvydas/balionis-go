package bago0

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	Name string `yaml:"name"`
	Timeout int `yaml:"timeout"`
}

type serverConfigWrapper struct {
	Server *ServerConfig `yaml:"server"`
}

func (c *ServerConfig) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t := serverConfigWrapper{ Server: c}

	if err := yaml.Unmarshal(data, &t); err != nil {
		return err
	}
	return nil
}
