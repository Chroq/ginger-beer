package repository_test

import (
	"database/sql"
	"ginger-beer/internal/app/adapter/repository"
	"ginger-beer/internal/app/domain"
	"ginger-beer/testdata/tu"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maxatome/go-testdeep/td"
)

func TestSqlRepository_GetComponents(t *testing.T) {
	type fields struct {
		DB func() *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.Component
		err    error
		clean  func()
	}{
		{
			name: "default",
			fields: fields{
				DB: func() *sql.DB {
					db, err := sql.Open("postgres", "postgresql://postgres@localhost/ginger-beer-test?sslmode=disable")
					td.CmpNoError(t, err)

					_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "test"(id SERIAL PRIMARY KEY, name VARCHAR(255));`)
					td.CmpNoError(t, err)

					return db
				},
			},
			want: &domain.Component{
				Schema: map[string]domain.Schema{
					"Test": {
						Type:       "object",
						Properties: map[string]domain.Property{"id": {Type: "integer"}, "name": {Type: "string", MaxLength: tu.Ptr(255)}},
					},
				},
			},
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
			r := &repository.SQLRepository{
				DB: tt.fields.DB(),
			}
			defer r.DB.Close()
			defer tt.clean()

			got, err := r.GetComponent()
			td.Cmp(t, err, tt.err)
			td.Cmp(t, got, tt.want)
		})
	}
}

func TestSQLRepository_GetEntities(t *testing.T) {
	type fields struct {
		DB func() *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
		error  error
		clean  func()
	}{
		{
			name: "default",
			fields: fields{
				DB: func() *sql.DB {
					db, err := sql.Open("postgres", "postgresql://postgres@localhost/ginger-beer-test?sslmode=disable")
					td.CmpNoError(t, err)

					_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "tests"(id SERIAL PRIMARY KEY, name VARCHAR(255));`)
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

				_, err = db.Exec(`DROP TABLE IF EXISTS "tests";`)
				td.CmpNoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.SQLRepository{
				DB: tt.fields.DB(),
			}
			got, err := r.GetEntities()
			td.Cmp(t, err, tt.error)
			td.Cmp(t, got, tt.want)
		})
	}
}
