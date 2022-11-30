package service

type PgTable struct {
	Name   string
	Fields []PgField
}

type PgField struct {
	Size *int
	Name string
	Type string
}
