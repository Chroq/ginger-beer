package service

import (
	"fmt"
	"os"
	"strings"
)

const (
	DefaultConnection           = "postgresql://localhost/postgres"
	DefaultDriver               = "postgres"
	DefaultFormat               = "json"
	DefaultOutput               = "."
	DefaultType                 = "basic"
	ErrMessageNotPermittedValue = "%s is not a permitted value for %s. Allowed values : %s"
	ErrValueSeparator           = ", "
	NameType                    = "type"
	NameFormat                  = "format"
	ParamConnection             = "-c"
	ParamFormat                 = "-f"
	ParamOutput                 = "-o"
	ParamType                   = "-t"
	ParamVersion                = "-v"
	PermittedFormatJSON         = "json"
	PermittedFormatYAML         = "yaml"
	PermittedTypeBasic          = "basic"
	PermittedTypeClean          = "clean"
	Version                     = "0.0.1"
)

type Config struct {
	Connection string
	Driver     string
	File       string
	Format     string
	Output     string
	Type       string
	Version    string
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

	if c.Connection != DefaultConnection {
		if !strings.HasPrefix(c.Connection, "postgresql://") {
			return fmt.Errorf("connection must be a postgresql connection")
		}
	}

	return nil
}

// NewConfig creates a new instance of the Config
func NewConfig() (*Config, error) {
	config := &Config{
		Output:     DefaultOutput,
		Driver:     DefaultDriver,
		Type:       DefaultType,
		Format:     DefaultFormat,
		Connection: DefaultConnection,
	}
	args := os.Args[1:]
	for i := range args {
		switch args[i] {
		case ParamVersion:
			config.Version = Version
			return config, nil
		case ParamFormat:
			config.Format = args[i+1]
		case ParamConnection:
			config.Connection = args[i+1]
		case ParamOutput:
			config.Output = args[i+1]
		case ParamType:
			config.Type = args[i+1]
		}
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}
