package main

import (
	"database/sql"
	"fmt"
)

type Sql struct {
	DriverName     string
	DataSourceName string
}

type Insert struct {
	sql          *Sql
	db           *sql.DB
	table        string
	columns      []string
	columnsValue map[string]interface{}
}

func main() {
	i := Insert{}
	i.columns = append(i.columns, `13`)
	fmt.Println(i)
}
