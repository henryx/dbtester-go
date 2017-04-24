/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"flag"
	"db"
)

func readCfg(filename string) *ini.File {
	res, err := ini.Load([]byte{}, filename)
	if err != nil {
		fmt.Println("Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func checkdb(engine string, sect *ini.Section) {
	var err error

	user := sect.Key("user").String()
	password := sect.Key("password").String()
	host := sect.Key("host").String()
	dbname := sect.Key("database").String()

	port, err := sect.Key("port").Int()
	if err != nil {
		fmt.Println("Malformed port value in configuration file:", err)
		os.Exit(1)
	}

	dbconn, err := db.OpenDB(engine, user, password, dbname, host, port)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		os.Exit(1)
	}
	defer dbconn.Close()

	if exist, err := db.CheckStructure(engine, dbname, dbconn); err != nil || !exist {
		err = db.CreateStructure(engine, dbconn)
		if err != nil {
			fmt.Println("Error crreating database structure:", err)
			os.Exit(1)
		}
	}
}

func main() {

	cfgfile := flag.String("cfg", "", "Set the configuration file")
	flag.Parse()

	cfg := readCfg(*cfgfile)
	sect := cfg.Section("mysql")
	engine := cfg.Section("general").Key("engine").String()

	checkdb(engine, sect)
}
