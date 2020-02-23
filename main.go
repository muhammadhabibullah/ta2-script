package main

import (
	"errors"
	"log"
	"os"

	broker "ta2-script/brokers"
	database "ta2-script/databases"

	"github.com/joho/godotenv"
)

func main() {
	cmdString := command()
	if cmdString == "raw" || cmdString == "finale" {
		godotenv.Load()
		database.InitMySQL()
		broker.InitMQTT(cmdString)
	} else {
		log.Fatal(errors.New("Command false"))
		return
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
