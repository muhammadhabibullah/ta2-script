package main

import (
	"os"

	broker "ta2-script/brokers"
	database "ta2-script/databases"
	logger "ta2-script/loggers"

	"github.com/joho/godotenv"
)

func main() {
	cmdString := command()
	if cmdString == "raw" || cmdString == "finale" {
		godotenv.Load()
		database.InitMySQL()
		broker.InitMQTT(cmdString)
	} else {
		logger.LogRedMessage("Command false:")
		logger.LogRedMessage("Run `go run main.go raw` or `go run main.go finale")
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
