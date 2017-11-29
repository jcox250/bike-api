package model

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

var DBType string
var ConnString string

func ConnectToDB() (*sql.DB, error) {
	return connectToDB()
}

func connectToDB() (*sql.DB, error) {
	conn, err := sql.Open(DBType, ConnString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return conn, nil
}

func SafeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && err != nil {
		*err = cerr
	}
}
