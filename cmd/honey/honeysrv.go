package main

import (
	"flag"
	"fmt"
	"honey/internal/config"
	"honey/internal/db"
	"honey/internal/service/honey"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := readConfig()

	db, err := db.NewDB(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := crateSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := honey.New(db, cfg)

	for _, s := range service.FindAll() {
		fmt.Println(s)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func crateSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS seasons (
		id integer primary key autoincrement,
		name varchar);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	insertSeason := `INSERT INTO seasons (name) VALUES (?)`
	s := fmt.Sprintf("Temporada numero %v", time.Now().Nanosecond())
	db.MustExec(insertSeason, s)
	return nil
}
