package service

import (
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
		env  func() []string
	}{
		{
			name: "default",
			want: &Config{
				File:   "model.go",
				Output: ".",
				Type:   "basic",
				Format: "json",
			},
		},
		{
			name: "full parameter",
			want: &Config{
				File:   "user.go",
				Output: "output",
				Type:   "clean",
				Format: "yaml",
			},
			env: func() []string {
				return []string{"main", "-f", "user.go", "-o", "output", "-t", "clean", "-format", "yaml"}
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != nil {
				os.Args = tt.env()
			}
			td.Cmp(t, NewConfig(), tt.want)
		})
	}
}
