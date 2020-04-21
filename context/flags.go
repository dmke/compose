package context

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

const (
	// ConfigFileName is the name of config file
	ConfigFileName = "config.json"
	configFileDir  = ".docker"
)

var (
	ConfigDir   string
	ContextName string
	ConfigFlag  = cli.StringFlag{
		Name:        "config",
		Usage:       "Location of client config files `DIRECTORY`",
		EnvVar:      "DOCKER_CONFIG",
		Value:       filepath.Join(home(), configFileDir),
		Destination: &ConfigDir,
	}

	ContextFlag = cli.StringFlag{
		Name:        "context, c",
		Usage:       "Name of the context `CONTEXT` to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with \"docker context use\")",
		EnvVar:      "DOCKER_CONTEXT",
		Destination: &ContextName,
	}
)

func home() string {
	home, _ := homedir.Dir()
	return home
}
