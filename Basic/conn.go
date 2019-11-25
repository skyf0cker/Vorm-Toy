package Basic

import (
	"database/sql"
	"time"

	//register driver
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	DB *sql.DB
	where string
}


func Connect(dsn string) (*Db, error) {

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return &Db{DB:conn}, nil
}

