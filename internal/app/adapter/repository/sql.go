package repository

import (
	"database/sql"
	"fmt"
)

type SqlRepository struct {
	DB *sql.DB
}

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Size *int
	Name string
	Type string
}

func (r *SqlRepository) GetTableNames() ([]Table, error) {
	var Tables []Table

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
		var tableName Table
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

func (r *SqlRepository) GetFields(tableName string) ([]Field, error) {
	var fields []Field

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
		var field Field
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
