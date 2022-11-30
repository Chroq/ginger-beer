package translator

import (
	"errors"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestSQLToOpenAPIScalarTypes(t *testing.T) {
	type args struct {
		sqlType string
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "bigint should return integer",
			args: args{
				sqlType: "bigint",
			},
			want: "integer",
		}, {
			name: "int8 should return integer",
			args: args{
				sqlType: "int8",
			},
			want: "integer",
		}, {
			name: "bigserial should return integer",
			args: args{
				sqlType: "bigserial",
			},
			want: "integer",
		}, {
			name: "serial8 should return integer",
			args: args{
				sqlType: "serial8",
			},
			want: "integer",
		}, {
			name: "bit should return string",
			args: args{
				sqlType: "bit",
			},
			want: "string",
		}, {
			name: "bit varying should return string",
			args: args{
				sqlType: "bit varying",
			},
			want: "string",
		}, {
			name: "varbit should return string",
			args: args{
				sqlType: "varbit",
			},
			want: "string",
		}, {
			name: "bool should return boolean",
			args: args{
				sqlType: "bool",
			},
			want: "boolean",
		}, {
			name: "boolean should return boolean",
			args: args{
				sqlType: "boolean",
			},
			want: "boolean",
		}, {
			name: "box should return string",
			args: args{
				sqlType: "box",
			},
			want: "string",
		}, {
			name: "bytea should return string",
			args: args{
				sqlType: "bytea",
			},
			want: "string",
		}, {
			name: "character should return string",
			args: args{
				sqlType: "character",
			},
			want: "string",
		}, {
			name: "character varying should return string",
			args: args{
				sqlType: "character varying",
			},
			want: "string",
		}, {
			name: "varchar should return string",
			args: args{
				sqlType: "varchar",
			},
			want: "string",
		}, {
			name: "cidr should return string",
			args: args{
				sqlType: "cidr",
			},
			want: "string",
		}, {
			name: "circle should return string",
			args: args{
				sqlType: "circle",
			},
			want: "string",
		}, {
			name: "date should return string",
			args: args{
				sqlType: "date",
			},
			want: "string",
		}, {
			name: "double precision should return number",
			args: args{
				sqlType: "double precision",
			},
			want: "number",
		}, {
			name: "inet should return string",
			args: args{
				sqlType: "inet",
			},
			want: "string",
		}, {
			name: "integer should return integer",
			args: args{
				sqlType: "integer",
			},
			want: "integer",
		}, {
			name: "int should return integer",
			args: args{
				sqlType: "int",
			},
			want: "integer",
		}, {
			name: "int4 should return integer",
			args: args{
				sqlType: "int4",
			},
			want: "integer",
		}, {
			name: "interval should return string",
			args: args{
				sqlType: "interval",
			},
			want: "string",
		}, {
			name: "json	should return string",
			args: args{
				sqlType: "json",
			},
			want: "string",
		}, {
			name: "jsonb should return string",
			args: args{
				sqlType: "jsonb",
			},
			want: "string",
		}, {
			name: "line should return string",
			args: args{
				sqlType: "line",
			},
			want: "string",
		}, {
			name: "lseg should return string",
			args: args{
				sqlType: "lseg",
			},
			want: "string",
		}, {
			name: "macaddr should return string",
			args: args{
				sqlType: "macaddr",
			},
			want: "string",
		}, {
			name: "macaddr8 should return string",
			args: args{
				sqlType: "macaddr8",
			},
			want: "string",
		}, {
			name: "money should return number",
			args: args{
				sqlType: "money",
			},
			want: "number",
		}, {
			name: "numeric should return number",
			args: args{
				sqlType: "numeric",
			},
			want: "number",
		}, {
			name: "decimal should return number",
			args: args{
				sqlType: "decimal",
			},
			want: "number",
		}, {
			name: "path should return string",
			args: args{
				sqlType: "path",
			},
			want: "string",
		}, {
			name: "pg_lsn should return string",
			args: args{
				sqlType: "pg_lsn",
			},
			want: "string",
		}, {
			name: "pg_snapshot should return string",
			args: args{
				sqlType: "pg_snapshot",
			},
			want: "string",
		}, {
			name: "point should return string",
			args: args{
				sqlType: "point",
			},
			want: "string",
		}, {
			name: "polygon should return string",
			args: args{
				sqlType: "polygon",
			},
			want: "string",
		}, {
			name: "real should return number",
			args: args{
				sqlType: "real",
			},
			want: "number",
		}, {
			name: "float4 should return number",
			args: args{
				sqlType: "float4",
			},
			want: "number",
		}, {
			name: "smallint should return integer",
			args: args{
				sqlType: "smallint",
			},
			want: "integer",
		}, {
			name: "int2 should return integer",
			args: args{
				sqlType: "int2",
			},
			want: "integer",
		}, {
			name: "smallserial should return integer",
			args: args{
				sqlType: "smallserial",
			},
			want: "integer",
		}, {
			name: "serial2 should return integer",
			args: args{
				sqlType: "serial2",
			},
			want: "integer",
		}, {
			name: "serial should return integer",
			args: args{
				sqlType: "serial",
			},
			want: "integer",
		}, {
			name: "serial4 should return integer",
			args: args{
				sqlType: "serial4",
			},
			want: "integer",
		}, {
			name: "text should return string",
			args: args{
				sqlType: "text",
			},
			want: "string",
		}, {
			name: "time should return string",
			args: args{
				sqlType: "time",
			},
			want: "string",
		}, {
			name: "timestamp should return string",
			args: args{
				sqlType: "timestamp",
			},
			want: "string",
		}, {
			name: "tsquery should return string",
			args: args{
				sqlType: "tsquery",
			},
			want: "string",
		}, {
			name: "tsvector should return string",
			args: args{
				sqlType: "tsvector",
			},
			want: "string",
		}, {
			name: "txid_snapshot should return string",
			args: args{
				sqlType: "txid_snapshot",
			},
			want: "string",
		}, {
			name: "uuid should return string",
			args: args{
				sqlType: "uuid",
			},
			want: "string",
		}, {
			name: "xml should return string",
			args: args{
				sqlType: "xml",
			},
			want: "string",
		}, {
			name: "invalid type should be detected",
			args: args{
				sqlType: "test",
			},
			want: "",
			err:  errors.New("invalid SQL type: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			openAPIType, err := PgSQLToOpenAPITypes(tt.args.sqlType)
			td.Cmp(t, err, tt.err)
			td.Cmp(t, openAPIType, tt.want)
		})
	}
}
