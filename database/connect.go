package database

import (
	"context"
    "log"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

func SQL_Connect(){
    
    

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