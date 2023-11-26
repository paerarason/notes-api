package api

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserCreateHandler() gin.HandlerFunc{
	return func (c *gin.Context){
	 db:=database.SQL_Connect()
	 //validate the Given Values with struct
     var user User
	 if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    
	if user.Username == "" || user.Password == "" || user.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username, password, and email are required"})
        return
    }

	return
	 user:=models.User{username:body.username,email:body.email,password:body.password,created_at:time.Now()}
	 return db.Create(&user) 
}
}

func UserGETHandler() gin.HandlerFunc{
	return func (c *gin.Context){
	 db:=database.SQL_Connect()
	 //validate the Given Values with struct
     var user User
	 if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    
	if user.Username == "" || user.Password == "" || user.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username, password, and email are required"})
        return
    }


	return
	 }
	 user:=models.User{username:body.username,email:body.email,password:body.password,created_at:time.Now()}
	 return db.Create(&user) 
}



