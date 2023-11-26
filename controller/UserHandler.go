package api

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/paerarason/notes-api/models"
)

func UserCreateHandler() gin.HandlerFunc{
	 return func (c *gin.Context){
	 db:=database.SQL_Connect()
	 
	 //validate the Given Values with struct
     var user models.User
	 if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    
	if user.Username == "" || user.Password == "" || user.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username, password, and email are required"})
        return
    }
    
	//hash the password to the database
	hash,herr:=HashPassword(body.password)
	if herr!=nil {
	 user:=models.User{username:body.username,email:body.email,password:body.password,created_at:time.Now()}
	 return db.Create(&user) }
	}
	
	c.JSON(http.StatusBadRequest, gin.H{"error": "Check the content Again "})
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
