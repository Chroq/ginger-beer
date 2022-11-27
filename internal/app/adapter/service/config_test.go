package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"syscall"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  *Config
		err   error
		env   func() []string
		clean func()
	}{
		{
			name: "default",
			want: &Config{
				Connection: "postgresql://localhost/postgres",
				Driver:     "postgres",
				Output:     ".",
				Type:       "basic",
				Format:     "json",
			},
			env: func() []string {
				return []string{"main"}
			},
		},
		{
			name: "full parameter",
			want: &Config{
				Connection: "postgresql://localhost/postgres",
				Driver:     "postgres",
				Output:     "output",
				Type:       "clean",
				Format:     "yaml",
			},
			env: func() []string {
				return []string{"main", "-c", "postgresql://localhost/postgres", "-o", "output", "-t", "clean", "-f", "yaml"}
			},
			clean: func() {
				if err := os.Remove("./output"); err != nil {
					td.CmpNoError(t, err)
				}
			},
		},
		{
			name: "only version",
			want: &Config{
				Connection: "postgresql://localhost/postgres",
				Driver:     "postgres",
				Output:     ".",
				Type:       "basic",
				Format:     "json",
				Version:    "0.0.1",
			},
			env: func() []string {
				return []string{"main", "-v"}
			},
		},
		{
			name: "wrong type",
			err:  fmt.Errorf("test is not a permitted value for type. Allowed values : basic, clean"),
			env: func() []string {
				return []string{"main", "-t", "test"}
			},
		},
		{
			name: "wrong format",
			err:  fmt.Errorf("test is not a permitted value for format. Allowed values : json, yaml"),
			env: func() []string {
				return []string{"main", "-f", "test"}
			},
		},
		{
			name: "existing directory",
			err: &fs.PathError{
				Op:   "mkdir",
				Path: "test",
				Err:  syscall.EEXIST,
			},
			env: func() []string {
				err := os.Mkdir("./test", os.ModePerm)
				td.CmpNoError(t, err)
				return []string{"main", "-o", "test"}
			},
			clean: func() {
				err := os.Remove("./test")
				td.CmpNoError(t, err)
			},
		},
		{
			name: "incorrectly formatted connection string",
			err:  errors.New("connection must be a postgresql connection"),
			env: func() []string {
				return []string{"main", "-c", "test"}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clean != nil {
				defer tt.clean()
			}

			if tt.env != nil {
				os.Args = tt.env()
			}

			config, err := NewConfig()
			td.Cmp(t, err, tt.err)
			td.Cmp(t, config, tt.want)
		})
	}
}
