package system

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataSource struct {
	DB *gorm.DB
}

type DbConfig struct {
	Host     string `json:"db_host"`
	Port     int    `json:"db_port"`
	Username string `json:"db_user"`
	Password string `json:"db_password"`
	Name     string `json:"db_name"`
}

func NewDbConfig() *DbConfig {
	var dbConfig DbConfig
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		panic("DB_HOST environment variable required but not set")
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		panic("DB_PORT environment variable required but not set")
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("DB_USER environment variable required but not set")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		panic("DB_PASSWORD environment variable required but not set")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		panic("DB_NAME environment variable required but not set")
	}
	dbConfig.Host = dbHost
	dbConfig.Port, _ = strconv.Atoi(dbPort)
	dbConfig.Username = dbUser
	dbConfig.Password = dbPassword
	dbConfig.Name = dbName
	return &dbConfig
}

func NewDataSource() *DataSource {
	var err error
	dbConfig := NewDbConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return &DataSource{DB: db}
}
