package main

import (
	"HomeAI/config"
	"HomeAI/http"
	"log"
	"os"
)

const usage string = "usage: ./HomeAI /path/to/config.json"

func main() {
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
