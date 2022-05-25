package mysql

import (
	"database/sql"
	"fmt"
)

type MySQL struct {
	*sql.DB
}

func (db *MySQL) insertRow(sql string) {
	_, err := db.DB.Exec(sql)
	if err != nil {
		fmt.Printf("exec err: %v\n", err)
		return
	}
}

func (db *MySQL) deleteRow(sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("exec err: %v\n", err)
		return
	}
}