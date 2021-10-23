package auth

import "github.com/golang-jwt/jwt"

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var userCredentials = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
