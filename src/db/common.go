/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package db

import (
	"database/sql"
	"errors"
	"strings"
	"fmt"
	"strconv"
)

func createViewTest() []string {
	var res []string
	for i := 1; i <= 10; i++ {
		res = append(res,
			fmt.Sprintf(
				strings.Join([]string{
					"CREATE TABLE viewtest%s(",
					"ins INTEGER,",
					"date_ins INTEGER,",
					"hour_ins INTEGER,",
					"txt VARCHAR(30))",
				}, " "), strconv.Itoa(i)))
	}

	res = append(res, strings.Join([]string{
		"CREATE OR REPLACE VIEW viewtest AS",
		"SELECT * FROM viewtest1 UNION",
		"SELECT * FROM viewtest2 UNION",
		"SELECT * FROM viewtest3 UNION",
		"SELECT * FROM viewtest4 UNION",
		"SELECT * FROM viewtest5 UNION",
		"SELECT * FROM viewtest6 UNION",
		"SELECT * FROM viewtest7 UNION",
		"SELECT * FROM viewtest8 UNION",
		"SELECT * FROM viewtest9 UNION",
		"SELECT * FROM viewtest10"},
		" "))

	return res
}

func OpenDB(engine, user, password, dbname, host string, port int) (*sql.DB, error) {
	var conn *sql.DB
	var err error

	switch engine {
	case "mysql":
		conn, err = openMySQLDB(user, password, host, dbname, port)
		return conn, err
	default:
		return nil, errors.New("Unknown database engine: " + engine)
	}
}

func CheckStructure(engine, dbname string, conn *sql.DB) (bool, error) {
	switch engine {
	case "mysql":
		return checkMySQLStructure(conn, dbname)
	default:
		return false, errors.New("Unknown database engine: " + engine)
	}
}

func CreateStructure(engine string, conn *sql.DB) error {
	switch engine {
	case "mysql":
		return createMySQLStructure(conn)
	default:
		return errors.New("Unknown database engine: " + engine)
	}
}
