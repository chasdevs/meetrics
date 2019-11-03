package data

import (
	"fmt"
	"github.com/chasdevs/meetrics/pkg/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func SetupDb() {
	// Get config
	setupConfig := conf.MysqlRootConfig()

	// Get connection
	db := getDb(&setupConfig)

	config := conf.MysqlConfig()

	// Make database
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v;", config.DBName)
	db.Exec(sql)

	// Create user
	sql = fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'%%' IDENTIFIED BY '%v';", config.User, config.Passwd)
	db.Exec(sql)

	sql = fmt.Sprintf("GRANT ALL ON %v.* TO '%v'@'%%';", config.DBName, config.User)
	db.Exec(sql)
}

func TeardownDb() {

	rootConfig := conf.MysqlRootConfig()
	db := getDb(&rootConfig)

	// Make database
	//db.Exec("DROP USER IF EXISTS meetrics;")
	db.Exec("DROP DATABASE IF EXISTS meetrics;")

}

func Migrate() {
	config := conf.MysqlConfig()
	db := getDb(&config)
	db.DropTableIfExists(&UserMeetingMins{}, &User{}, &Meeting{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &UserMeetingMins{}, &Meeting{})
	db.Model(&UserMeetingMins{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func getDb(config *mysql.Config) *gorm.DB {
	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db
}
