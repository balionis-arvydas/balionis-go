package bago0

import (
	"github.com/balionis-arvydas/balionis-go/bago0/pkg/logutil"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	Name   string               `yaml:"name"`
	Logger logutil.LoggerConfig `yaml:"loggerConfig"`
}

type serverConfigWrapper struct {
	Server *ServerConfig `yaml:"server"`
}

func (c *ServerConfig) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read "+filename)
	}
	t := serverConfigWrapper{Server: c}

	if err := yaml.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "failed to parse "+filename)
	}
	return nil
}
