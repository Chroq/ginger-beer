package repository

import (
	"database/sql"
	"fmt"
)

type SqlRepository struct {
	DB *sql.DB
}

func (r *SqlRepository) GetTableNames() ([]string, error) {
	var tableNames []string

	query, err := r.DB.Query(`
		select table_name 
		from information_schema.tables
		where table_schema != 'pg_catalog' and table_schema != 'information_schema'
		`)
	if err != nil {
		return tableNames, err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}(query)

	for query.Next() {
		var tableName string
		if err := query.Scan(&tableName); err != nil {
			return tableNames, err
		}
		tableNames = append(tableNames, tableName)
	}
	if err = query.Err(); err != nil {
		return tableNames, err
	}

	return tableNames, nil
}
