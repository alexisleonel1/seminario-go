package main

import (
	"flag"
	"fmt"
	"honey/internal/config"
	"honey/internal/service/honey"
	"os"
)

func main() {

	cfg := readConfig()

	service, _ := honey.New(cfg)
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
