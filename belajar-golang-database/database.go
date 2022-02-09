package belajargolangdatabase

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/localdata_go?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)  //dlm 5 menit tidak bekerja di stop
	db.SetConnMaxLifetime(60 * time.Minute) //jika dlm 60 menit koneksi akan diperaharui

	return db
}
