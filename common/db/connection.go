package db

import (
	// "common/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // this is needed to use mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DbUser     = "root"
	DbPassword = "admin"
	DbHost     = "localhost"
	DbPort     = "3306"
	DbName     = "albumsDB"
)

func SetupDatabaseConn() (*gorm.DB, error) {
	// ?charset=utf8mb4 is for support for storing any unicode character
	// parseTime = true is to convert MySQL date and datetime types to Go's time.Time type
	// loc=Local sets the timezone to the one the server is run in
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort)
	// dbConn, err := initiateConnection(connectionString)
	// if err != nil {
	// 	return nil, err
	// }

	// err = dbConn.Exec("CREATE DATABASE IF NOT EXISTS " + DbName).Error
	// if err != nil {
	// 	return nil, err
	// }

	// sqlDB, err := dbConn.DB()
	// if err != nil {
	// 	return nil, err
	// }
	// defer sqlDB.Close() // close connection when SetupDatabaseConn ends

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	dbConn, err := initiateConnection(connectionString)
	if err != nil {
		return nil, err
	}

	// err = dbConn.AutoMigrate(&models.Album{}, &models.AlbumType{})
	// if err != nil {
	// 	return nil, err
	// }

	return dbConn, nil
}

func initiateConnection(connString string) (*gorm.DB, error) {
	DbConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	return DbConn, err
}
