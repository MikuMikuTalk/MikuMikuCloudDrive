package jwts

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
)

type JwtPayload struct {
	UserID   uint   `json:"userID"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

// 构建GTI
func generateJTI() string {
	return uuid.NewString()
}

func ExtractJTI(claims *CustomClaims) string {
	return claims.ID
}

func GenerateJwtToken(payload JwtPayload, accessSecret string, expires int) (string, error) {
	claims := CustomClaims{
		JwtPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        generateJTI(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}
	return "Bearer " + signedToken, nil // 添加 Bearer 前缀
}

func ParseJwtToken(jwtToken string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法jwt token")
}
func ProcessJwtToken(jwtToken string) string {
	token := strings.TrimPrefix(jwtToken, "Bearer ")
	return token
}

// func ValidateJwtToken(jwtToken string, accessSecret string) (bool, error) {
// 	_, err := ParseJwtToken(jwtToken, accessSecret)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
