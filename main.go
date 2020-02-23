package main

import (
	"os"

	broker "ta2-script/brokers"
	database "ta2-script/databases"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.InitMySQL()
	cmdString := command()
	broker.InitMQTT(cmdString)
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
