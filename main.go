package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


type Message struct {
	Sender string `json:"sender"`
	Content string `json:"content"`
}

func main(){
	r:=gin.Default()


	//endpoint 
	r.GET("/ws", func(c *gin.Context){
		header := make(http.Header)
		    conn, err := websocket.Upgrade(c.Writer, c.Request, header, 1024, 1024)
			if err!=nil{
				log.Fatal(err)
			}
			
			go handleMessage(conn)
			for {
				_,_ ,err:=conn.ReadMessage()
				if err!=nil {
					log.Println(err)
					break
				}
			}  
			conn.Close()
   })
   r.Run(":8080")

  
}

func handleMessage(conn *websocket.Conn) {
    for {
        // Read the incoming message
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            break
        }

        // Unmarshal the JSON message
        var msg Message
        err = json.Unmarshal(message, &msg)
        if err != nil {
            log.Println(err)
            continue
        }

        // Broadcast the message to all connected clients
        fmt.Println(msg.Sender, msg.Content)
    }
}
