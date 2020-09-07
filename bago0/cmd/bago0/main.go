package main

import (
	"flag"
	"fmt"
	"github.com/balionis-arvydas/balionis-go/bago0/pkg/bago0"
	"github.com/balionis-arvydas/balionis-go/bago0/pkg/logutil"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	server     bago0.ServerConfig
	configFile string
}

func (c *Config) Setup() error {
	flag.StringVar(&c.configFile, "configFile", "default.yaml", "yaml file to load")
	flag.Parse()

	if err := c.server.Load(c.configFile); err != nil {
		return errors.Wrap(err, "failed to load configFile "+c.configFile)
	}
	return nil
}

func main() {
	var config Config
	if err := config.Setup(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to setup: %v\n", err)
		os.Exit(1)
	}

	logger, err := logutil.NewLogger(config.server.Logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to create logger: %v\n", err)
		os.Exit(1)
	}

	level.Info(logger).Log("config", fmt.Sprintf("%+v", config))

	level.Info(logger).Log("done", "+")
}
