package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // this is needed to use mysql

)

const (
	DbUser     = "root"
	DbPassword = "admin"
	DbHost     = "localhost"
	DbPort     = "3306"
	DbName     = "albumsDB"
)

var DbConn *sql.DB

func SetupDatabaseConn() error {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbName)
	DbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	err = DbConn.Ping()
	if err != nil {
		return err
	}

	createAlbumsTableQuery := `CREATE TABLE IF NOT EXISTS Album (
		id VARCHAR(255) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		artist VARCHAR(255) NOT NULL,
		price DOUBLE NOT NULL
	);`

	_, err = DbConn.Exec(createAlbumsTableQuery)
	if err != nil {
		return err
	}

	return nil
}