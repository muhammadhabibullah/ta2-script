package database

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

// DB :nodoc:
var DB *gorm.DB

// InitMySQL database
func InitMySQL() (*gorm.DB, error) {
	var (
		USERNAME = os.Getenv("DB_USERNAME")
		PASSWORD = os.Getenv("DB_PASSWORD")
		HOST     = os.Getenv("DB_HOST")
		PORT     = os.Getenv("DB_PORT")
		DATABASE = os.Getenv("DB_DATABASE")

		mysqlCon = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			USERNAME,
			PASSWORD,
			HOST,
			PORT,
			DATABASE,
		)
	)
	var err error
	DB, err = gorm.Open("mysql", mysqlCon)
	if err != nil {

		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", mysqlCon))
		errorOutput.Println(fmt.Errorf("%s", err))

	} else {

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)

		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", mysqlCon))

	}

	return DB, err
}
