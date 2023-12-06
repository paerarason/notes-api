package users

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func UserCreateHandler() gin.HandlerFunc{
   return func (c *gin.Context){
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
