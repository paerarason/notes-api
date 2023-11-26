package database
import (
    "log"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func SQL_Connect() (*gorm.DB){
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err!=nil{
        log.Fatal(err)
    }
    return db 
}

