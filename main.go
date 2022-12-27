package main

import (
	"flag"
	"github.com/MagonxESP/dropper/http"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	envFile := flag.String("env-file", ".env", "The .env file location")
	flag.Parse()

	err := godotenv.Load(*envFile)

	if err != nil {
		log.Println(err)
	}

	http.StartHttpServer()
}
