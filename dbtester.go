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

type ConnData struct {
	Engine   string
	User     string
	Password string
	Dbname   string
	Host     string
	Port     int
}

func readCfg(filename string) *ini.File {
	res, err := ini.Load([]byte{}, filename)
	if err != nil {
		fmt.Println("Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func createdb(dbdata ConnData) {
	var err error

	dbconn, err := db.OpenDB(dbdata.Engine, dbdata.User, dbdata.Password, dbdata.Dbname, dbdata.Host, dbdata.Port)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		os.Exit(1)
	}
	defer dbconn.Close()

	if exist, err := db.CheckStructure(dbdata.Engine, dbdata.Dbname, dbconn); err != nil || !exist {
		err = db.CreateStructure(dbdata.Engine, dbconn)
		if err != nil {
			fmt.Println("Error crreating database structure:", err)
			os.Exit(1)
		}
	}
}

func connInfo(engine string, sect *ini.Section) ConnData {
	var res ConnData
	var err error

	res.Engine = engine
	res.User = sect.Key("user").String()
	res.Password = sect.Key("password").String()
	res.Host = sect.Key("host").String()
	res.Dbname = sect.Key("database").String()

	res.Port, err = sect.Key("port").Int()
	if err != nil {
		fmt.Println("Malformed port value in configuration file:", err)
		os.Exit(1)
	}

	return res
}

func main() {

	cfgfile := flag.String("cfg", "", "Set the configuration file")
	flag.Parse()

	if *cfgfile == "" {
		fmt.Println("No configuration file passed. See help")
		os.Exit(1)
	}

	cfg := readCfg(*cfgfile)
	engine := cfg.Section("general").Key("engine").String()
	sect := cfg.Section(engine)

	dbinfo := connInfo(engine, sect)
	createdb(dbinfo)
}
