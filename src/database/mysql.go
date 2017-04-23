/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)



func OpenDB(user, password, host, dbname string, port int) (*sql.DB, error) {
	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}