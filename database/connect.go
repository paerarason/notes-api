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

// func Db_connect()  *mongo.Client{
//      err:=godotenv.Load(".env")
//      if err!=nil{
//         log.Fatal(err)
//      }
//     clientOptions := options.Client().ApplyURI(os.Getenv("mongo_url"))
//     client, err := mongo.Connect(context.Background(), clientOptions)
//     if err != nil {
//         log.Fatal(err)
//     }

//     err = client.Ping(context.Background(), nil)
//     if err != nil {
//         log.Fatal(err)
//     }
    
// 	return client 
// }