package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type JWTClaim struct {
	UserId   uuid.UUID
	UserType user.Type
	jwt.StandardClaims
}

func GenerateJWT(userId uuid.UUID, userType user.Type) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := JWTClaim{
		UserId:   userId,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.UnixMilli(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

}
