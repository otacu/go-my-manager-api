package jwt_config

import "github.com/dgrijalva/jwt-go"

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

const (
	SECRET = "the_secret_for_sign"
)
