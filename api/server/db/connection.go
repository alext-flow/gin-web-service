package db

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql" // this is needed to use mysql
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "common/models"
)

const (
	DbUser     = "root"
	DbPassword = "admin"
	DbHost     = "localhost"
	DbPort     = "3306"
	DbName     = "albumsDB"
)

var DbConn *gorm.DB

func SetupDatabaseConn() error {
    // ?charset=utf8mb4 is for support for storing any unicode character
    // parseTime = true is to convert MySQL date and datetime types to Go's time.Time type
    // loc=Local sets the timezone to the one the server is run in
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort)
    err := initiateConnection(connectionString)
    if err != nil {
        return err
    }

    err = DbConn.Exec("CREATE DATABASE IF NOT EXISTS " + DbName).Error
    if err != nil {
        return err
    }

    sqlDB, err := DbConn.DB()
    if err != nil {
        return err
    }
    defer sqlDB.Close()  // close connection when SetupDatabaseConn ends

    connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
    err = initiateConnection(connectionString)
    if err != nil {
        return err
    }

    return DbConn.AutoMigrate(&models.Album{}, &models.AlbumType{})
}

func initiateConnection(connString string) error {
    var err error
    DbConn, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
    return err
}