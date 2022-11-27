package repository

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maxatome/go-testdeep/td"
)

func TestSqlRepository_GetTableNames(t *testing.T) {
	type fields struct {
		DB func() *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
		err    error
		clean  func()
	}{
		{
			name: "default",
			fields: fields{
				DB: func() *sql.DB {
					db, err := sql.Open("postgres", "postgresql://postgres@localhost/ginger-beer-test?sslmode=disable")
					td.CmpNoError(t, err)

					_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "test"(id SERIAL PRIMARY KEY);`)
					td.CmpNoError(t, err)

					return db
				},
			},
			want: []string{"test"},
			clean: func() {
				db, err := sql.Open("postgres", "postgresql://postgres@localhost/ginger-beer-test?sslmode=disable")
				defer func(db *sql.DB) {
					err := db.Close()
					td.CmpNoError(t, err)
				}(db)
				td.CmpNoError(t, err)

				_, err = db.Exec(`DROP TABLE IF EXISTS "test";`)
				td.CmpNoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SqlRepository{
				DB: tt.fields.DB(),
			}
			defer r.DB.Close()
			defer tt.clean()

			got, err := r.GetTableNames()
			td.Cmp(t, err, tt.err)
			td.Cmp(t, got, tt.want)
		})
	}
}
