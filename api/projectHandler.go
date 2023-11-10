package api

import (
	"github.com/paerarason/notes-api/database"
)

func Projecthandler(){
	client:=database.Db_connect()
	defer client.Close()	
}

