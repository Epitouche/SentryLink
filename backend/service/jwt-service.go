package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string, name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserIdfromJWTToken(tokenString string) (userId uint64, err error)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "email@example.com",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	return secret
}

func (jwtSrv *jwtService) GenerateToken(userId string, username string, admin bool) string {

	// Set custom and standard claims
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
			Id:        userId,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}

func (jwtSrv *jwtService) GetUserIdfromJWTToken(tokenString string) (userId uint64, err error) {

	token, err := jwtSrv.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		if jti, ok := claims["jti"].(string); ok {
			id, err := strconv.ParseUint(jti, 10, 64)
			if err != nil {
				return 0, err
			}
			return id, nil
		}
		return 0, fmt.Errorf("jti claim is not a float64")
	} else {
		return 0, err
	}
}
