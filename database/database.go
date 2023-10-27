package database

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielDDHM/world-coin-converter/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func StartDatabase() {
	DbHost := config.GetConfig().DatabaseHost
	DbPort := config.GetConfig().DatabasePort
	DbUser := config.GetConfig().DatabaseUser
	DbName := config.GetConfig().DatabaseName
	// DbSSlMode := config.GetConfig().DatabaseSslMode
	DbPass := config.GetConfig().DatabasePass
	DbMaxIddleConns := config.GetConfig().DatabaseMaxIdleConns
	DbMaxOpensConns := config.GetConfig().DatabaseMaxOpensConns

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		fmt.Println("Could not connect to the MySql Database")
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := database.DB()

	config.SetMaxIdleConns(DbMaxIddleConns)
	config.SetMaxOpenConns(DbMaxOpensConns)
	config.SetConnMaxLifetime(time.Hour)

	// migrations.RunAutoMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
