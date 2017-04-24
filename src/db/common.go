/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package db

import "database/sql"

func checkMySQLStructure(db *sql.DB) bool {
	return true
}

func CheckStructure(db *sql.DB) bool {
	return checkMySQLStructure(db)
}
