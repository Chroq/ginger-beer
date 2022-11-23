package service

import (
	"os"
)

const (
	DefaultFile   = "model.go"
	DefaultOutput = "."
	DefaultType   = "basic"
	DefaultFormat = "json"
	ParamFile     = "-f"
	ParamOutput   = "-o"
	ParamType     = "-t"
	ParamFormat   = "-format"
	ParamVersion  = "-v"
	Version       = "0.0.1"
)

type Config struct {
	Version string
	File    string
	Output  string
	Type    string
	Format  string
}

// NewConfig creates a new instance of the Config
func NewConfig() *Config {
	config := &Config{
		File:   DefaultFile,
		Output: DefaultOutput,
		Type:   DefaultType,
		Format: DefaultFormat,
	}
	args := os.Args[1:]
	for i := range args {
		switch args[i] {
		case ParamVersion:
			config.Version = Version
			break
		case ParamFile:
			config.File = args[i+1]
			i++
		case ParamOutput:
			config.Output = args[i+1]
			i++
		case ParamType:
			config.Type = args[i+1]
			i++
		case ParamFormat:
			config.Format = args[i+1]
			i++
		}
	}

	return config
}
