package main

import (
	"fmt"
	"os"

	broker "ta2-script/brokers"
	database "ta2-script/databases"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	cmdString := command()
	if cmdString == "raw" || cmdString == "finale" {
		godotenv.Load()
		database.InitMySQL()
		broker.InitMQTT(cmdString)
	} else {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println(fmt.Sprintf("Command false:"))
		errorOutput.Println(fmt.Sprintf("Run `go run main.go raw` or `go run main.go finale`"))
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
