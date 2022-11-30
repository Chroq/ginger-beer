package factory

import (
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/domain"
	"ginger-beer/testdata/tu"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestBuildPropertiesByPgTable(t *testing.T) {
	type args struct {
		table service.PgTable
	}
	tests := []struct {
		name string
		args args
		want map[string]domain.Property
		err  error
	}{
		{
			name: "default",
			args: args{
				table: service.PgTable{
					Name: "test",
					Fields: []service.PgField{
						{
							Name: "test",
							Type: "character varying",
							Size: tu.Ptr(255),
						},
					},
				},
			},
			want: map[string]domain.Property{
				"test": {
					Type:      "string",
					MaxLength: tu.Ptr(255),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildPropertiesByPgTable(tt.args.table)
			td.Cmp(t, err, tt.err)
			td.Cmp(t, got, tt.want)
		})
	}
}

func TestBuildSchemaByPgTable(t *testing.T) {
	type args struct {
		table service.PgTable
	}
	tests := []struct {
		name string
		args args
		want *domain.Schema
		err  error
	}{
		{
			name: "default",
			args: args{
				table: service.PgTable{
					Name: "test",
					Fields: []service.PgField{
						{
							Name: "test",
							Type: "character varying",
							Size: tu.Ptr(255),
						},
					},
				},
			},
			want: &domain.Schema{
				Type: "object",
				Properties: map[string]domain.Property{
					"test": {
						Type:      "string",
						MaxLength: tu.Ptr(255),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildSchemaByPgTable(tt.args.table)
			td.Cmp(t, err, tt.err)
			td.Cmp(t, got, tt.want)
		})
	}
}
