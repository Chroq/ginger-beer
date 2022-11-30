package domain

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGetOutputSchemaReference(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return a valid reference",
			args: args{
				name: "test",
			},
			want: "#/components/schemas/output.Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td.Cmp(t, GetOutputSchemaReference(tt.args.name), tt.want)
		})
	}
}
