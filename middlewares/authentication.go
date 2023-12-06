package middlewares
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paerarason/notes-api/database"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"
	//"fmt"
	//"os"

)

type Claims struct {
	AccountID int `json:"account_id"`
	jwt.StandardClaims
}
func JWTokenMiddlerware(c *gin.Context){

	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("authkey"), nil // Use your secret key here
        })

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
         c.Set("account_id", claims["account_id"])
	     c.Next()
	 }

	
}



func GenerateToken() gin.HandlerFunc {
    return func(c *gin.Context){
	username := c.PostForm("username")
	//password := c.PostForm("password")
	
	db,err:=database.SQL_Connect()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
			return
	}

    //get the account password from the database 
	var hash string 
	var  accountID int
	Query:=`SELECT password,ID FROM account WHERE account.username=$1` 
    
	err=db.QueryRow(Query,username).Scan(&hash,&accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get Account details"})
		return
	}
    
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := jwt.MapClaims{
        "account_id": accountID,
		"exp":expirationTime.Unix(),
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString([]byte("authkey"))
	if err != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}


func CheckPasswordHash(password string , hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}