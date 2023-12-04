package database
import (
    "log"
    "database/sql"
)

func SQL_Connect() *sql.DB{
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
    db, err :=sql.Open("postgres",dsn)
    if err!=nil{
        log.Fatal(err)
    } 
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    return db 
}

