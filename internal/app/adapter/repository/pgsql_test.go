package repository_test

import (
	"database/sql"
	"ginger-beer/internal/app/adapter/repository"
	"ginger-beer/internal/app/domain/valueobject"
	"ginger-beer/testdata/tu"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maxatome/go-testdeep/td"
)

func TestSQLRepository_GetEntities(t *testing.T) {
	type fields struct {
		DB func() *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string][]*valueobject.Field
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
			want: map[string][]*valueobject.Field{
				"test": {
					{
						Name: "id",
						Type: "integer",
					},
					{
						Name:      "name",
						Type:      "string",
						MaxLength: tu.Ptr(255),
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
			defer tt.clean()
			got, err := r.GetEntities()
			td.Cmp(t, err, tt.error)
			td.Cmp(t, got, tt.want)
		})
	}
}
