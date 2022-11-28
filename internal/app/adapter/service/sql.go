package service

type SQLTable struct {
	Name   string
	Fields []SQLField
}

type SQLField struct {
	Size *int
	Name string
	Type string
}
