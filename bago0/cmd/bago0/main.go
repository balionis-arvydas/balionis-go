package main

import (
	"flag"
	"fmt"
	"github.com/balionis-arvydas/balionis-go/bago0/pkg/bago0"
	"os"
)

type Config struct {
	server bago0.ServerConfig
	printConfig  bool
	configFile   string
}

func (c *Config) Setup() error {
	flag.BoolVar(&c.printConfig, "print.config", false, "Print the entire config")
	flag.StringVar(&c.configFile, "config.file", "default.yaml", "yaml file to load")
	flag.Parse()

	return c.server.Load(c.configFile)
}

func main() {
	var config Config
	if err := config.Setup(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to setup: %v\n", err)
		os.Exit(1)
	}

	if config.printConfig {
		fmt.Fprintf(os.Stdout, "DEBUG: config: %v\n", config)
	}

	fmt.Println("DEBUG: done")
}