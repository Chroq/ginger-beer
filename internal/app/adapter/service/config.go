package service

import (
	"fmt"
	"os"
	"strings"
)

const (
	PermittedFormatJSON         = "json"
	PermittedFormatYAML         = "yaml"
	PermittedTypeBasic          = "basic"
	PermittedTypeClean          = "clean"
	DefaultFile                 = "model.go"
	DefaultOutput               = "."
	DefaultType                 = "basic"
	DefaultFormat               = "json"
	ErrMessageNotPermittedValue = "%s is not a permitted value for %s. Allowed values : %s"
	ErrValueSeparator           = ", "
	NameType                    = "type"
	NameFormat                  = "format"
	ParamFile                   = "-f"
	ParamOutput                 = "-o"
	ParamType                   = "-t"
	ParamFormat                 = "-format"
	ParamVersion                = "-v"
	Version                     = "0.0.1"
)

type Config struct {
	Version string
	File    string
	Output  string
	Type    string
	Format  string
}

func (c *Config) Validate() error {
	if c.Type != PermittedTypeBasic && c.Type != PermittedTypeClean {
		permittedTypes := strings.Join([]string{PermittedTypeBasic, PermittedTypeClean}, ErrValueSeparator)
		return fmt.Errorf(ErrMessageNotPermittedValue, c.Type, NameType, permittedTypes)
	}

	if c.Format != PermittedFormatYAML && c.Format != PermittedFormatJSON {
		permittedFormats := strings.Join([]string{PermittedFormatJSON, PermittedFormatYAML}, ErrValueSeparator)
		return fmt.Errorf(ErrMessageNotPermittedValue, c.Format, NameFormat, permittedFormats)
	}

	if c.Output != "." {
		if err := os.Mkdir(c.Output, os.ModePerm); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if _, err := os.Stat(c.File); os.IsNotExist(err) {
		return err
	}

	return nil
}

// NewConfig creates a new instance of the Config
func NewConfig() (*Config, error) {
	config := buildConfig()

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// buildConfig creates a new instance of the Config
func buildConfig() *Config {
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
