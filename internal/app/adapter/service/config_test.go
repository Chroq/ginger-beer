package service

import (
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
				File:   "model.go",
				Output: ".",
				Type:   "basic",
				Format: "json",
			},
			env: func() []string {
				err := os.WriteFile("./model.go", []byte{}, os.ModePerm)
				td.CmpNoError(t, err)
				return []string{"main"}
			},
			clean: func() {
				err := os.Remove("./model.go")
				td.CmpNoError(t, err)
			},
		},
		{
			name: "full parameter",
			want: &Config{
				File:   "testModel.go",
				Output: "output",
				Type:   "clean",
				Format: "yaml",
			},
			env: func() []string {
				err := os.WriteFile("./testModel.go", []byte{}, os.ModePerm)
				td.CmpNoError(t, err)
				return []string{"main", "-f", "testModel.go", "-o", "output", "-t", "clean", "-format", "yaml"}
			},
			clean: func() {
				err := os.Remove("./testModel.go")
				td.CmpNoError(t, err)

				if err := os.Remove("./output"); err != nil {
					td.CmpNoError(t, err)
				}
			},
		},
		{
			name: "only version",
			want: &Config{
				Version: "0.0.1",
				File:    "model.go",
				Output:  ".",
				Type:    "basic",
				Format:  "json",
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
				return []string{"main", "-format", "test"}
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
		}, {
			name: "not existing file",
			err: &fs.PathError{
				Op:   "stat",
				Path: "testModel.go",
				Err:  syscall.Errno(0x2),
			},
			env: func() []string {
				err := os.WriteFile("../../../../testdata/testModel.go", []byte{}, os.ModePerm)

				td.CmpNoError(t, err)
				return []string{"main", "-f", "testModel.go"}
			},
			clean: func() {
				err := os.Remove("../../../../testdata/testModel.go")
				td.CmpNoError(t, err)
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
