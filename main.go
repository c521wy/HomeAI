package main

import (
	"HomeAI/config"
	"HomeAI/http"
	"HomeAI/version"
	"log"
	"os"
	"time"
)

const usage string = "usage: ./HomeAI /path/to/config.json"

func main() {
	log.Println("git version: ", version.GitVersion)
	log.Println("build time: ", version.BuildTime)
	log.Println("start time: ", time.Now())

	if len(os.Args) < 2 {
		log.Fatalln(usage)
	}
	err := config.Load(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	err = http.ServerRun(config.AppConfig.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
