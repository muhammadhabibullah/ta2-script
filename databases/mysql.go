package database

import (
	"fmt"
	"os"

	logger "ta2-script/loggers"

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
		logger.LogRedError(err)
		logger.LogRedMessage("Failed connected to database %s", mysqlCon)
	} else {
		logger.LogGreenMessage("Successfully connected to database %s", mysqlCon)
	}

	return DB, err
}
