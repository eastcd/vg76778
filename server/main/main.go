package main

import (
	"flag"
	"github.com/eastcd/vg76778/service"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "../config.yaml", "config file path")
	flag.Parse()
}

func main() {
	svc := service.NewService(configPath)
	if err := svc.Start(); err != nil {
		log.Fatalln(err)
	}
}
