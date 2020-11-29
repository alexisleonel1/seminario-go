package main

import (
	"flag"
	"fmt"
	"honey/internal/config"
	"honey/internal/db"
	"honey/internal/service/honey"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := readConfig()

	db, err := db.NewDB(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := honey.New(db, cfg)
	httpService := honey.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
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
