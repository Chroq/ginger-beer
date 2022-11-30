package repository

import (
	"database/sql"
	"fmt"
	"ginger-beer/internal/app/adapter/factory"
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/domain/valueobject"

	"github.com/tangzero/inflector"
)

const (
	PgQueryGetTables = `
		select table_name 
		from information_schema.tables
		where table_schema != 'pg_catalog' and table_schema != 'information_schema'
		`
	PqQueryGetFields = `
			SELECT column_name, data_type, character_maximum_length
			FROM information_schema."columns" c
			WHERE TABLE_NAME = '%s'
		`
)

type SQLRepository struct {
	DB *sql.DB
}

func (r *SQLRepository) GetEntities() (map[string][]*valueobject.Field, error) {
	pgTables, err := r.getTableNames()
	if err != nil {
		return nil, err
	}
	entities := make(map[string][]*valueobject.Field, len(pgTables))
	for i := range pgTables {
		entity := inflector.Singularize(pgTables[i].Name)
		fields, err := r.getFieldsByEntity(entity)
		if err != nil {
			return nil, err
		}
		entities[entity] = fields
	}

	return entities, nil
}

func (r *SQLRepository) getFieldsByEntity(entity string) ([]*valueobject.Field, error) {
	var pgFields []service.PgField
	queryString := fmt.Sprintf(PqQueryGetFields, inflector.Pluralize(entity))
	query, err := r.DB.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}(query)

	for query.Next() {
		var field service.PgField
		if err := query.Scan(&field.Name, &field.Type, &field.Size); err != nil {
			return nil, err
		}
		pgFields = append(pgFields, field)
	}
	if err = query.Err(); err != nil {
		return nil, err
	}

	if fields, err := factory.PgFieldsToFields(pgFields); err != nil {
		return nil, err
	} else {
		return fields, nil
	}
}

func (r *SQLRepository) getTableNames() ([]service.PgTable, error) {
	var Tables []service.PgTable

	query, err := r.DB.Query(PgQueryGetTables)
	if err != nil {
		return Tables, err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}(query)

	for query.Next() {
		var tableName service.PgTable
		if err := query.Scan(&tableName.Name); err != nil {
			return Tables, err
		}
		Tables = append(Tables, tableName)
	}
	if err = query.Err(); err != nil {
		return Tables, err
	}

	return Tables, nil
}
