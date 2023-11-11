package api

import (
	"context"
	"time"
	"github.com/paerarason/notes-api/database"
)

func Projecthandler(username string){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client:=database.Db_connect()
	defer client.Disconnect(ctx)
	collection:=client.Database("kanban").Collection("projects")
	collection.find({"user":username})
	return 
}

