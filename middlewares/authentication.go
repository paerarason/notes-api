package middlewares


func JWTAuthentication(){
  mySigningKey := []byte("AllYourBase")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

// Create the Claims
claims := MyCustomClaims{
	"bar",
	jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	},
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
ss, err := token.SignedString(mySigningKey)     
}