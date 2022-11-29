package repository

import (
	"database/sql"
	"fmt"
	"ginger-beer/internal/app/adapter/factory"
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/domain"
)

type SQLRepository struct {
	DB *sql.DB
}

func (r *SQLRepository) GetComponent() (*domain.Component, error) {
	var component domain.Component
	tables, err := r.getTableNames()
	if err != nil {
		return nil, nil
	}

	component.Schema = make(map[string]domain.Schema, len(tables))
	for i := range tables {
		tables[i].Fields, err = r.getFields(tables[i].Name)
		if err != nil {
			return nil, nil
		}
		schema, err := factory.BuildSchemaBySQLTable(tables[i])
		if err != nil {
			return nil, err
		}
		component.Schema[tables[i].Name] = *schema
	}

	return &component, nil
}

func (r *SQLRepository) getTableNames() ([]service.SQLTable, error) {
	var Tables []service.SQLTable

	query, err := r.DB.Query(`
		select table_name 
		from information_schema.tables
		where table_schema != 'pg_catalog' and table_schema != 'information_schema'
		`)
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
		var tableName service.SQLTable
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

func (r *SQLRepository) getFields(tableName string) ([]service.SQLField, error) {
	var fields []service.SQLField

	queryString := fmt.Sprintf(`
			SELECT column_name, data_type, character_maximum_length
			FROM information_schema."columns" c
			WHERE TABLE_NAME = '%s'
		`, tableName)
	query, err := r.DB.Query(queryString)
	if err != nil {
		return fields, err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}(query)

	for query.Next() {
		var field service.SQLField
		if err := query.Scan(&field.Name, &field.Type, &field.Size); err != nil {
			return fields, err
		}
		fields = append(fields, field)
	}
	if err = query.Err(); err != nil {
		return fields, err
	}

	return fields, nil
}
