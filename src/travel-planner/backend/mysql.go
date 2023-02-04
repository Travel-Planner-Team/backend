package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *MySQLBackend
)

type MySQLBackend struct {
	db *gorm.DB
}

func InitMySQLBackend(config *util.MySQLInfo) {

	endpoint, username, password := config.Endpoint, config.Username, config.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		username, password, endpoint, constants.MYSQL_DBNAME, constants.MYSQL_CONFIG)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = &MySQLBackend{db}
}
