package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

type TokenMetaData struct {
	Expires int64
}

func GenerateNewAccessToken() (string, error) {

	secret := GetValue("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil

}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetaData, error) {

	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))

		return &TokenMetaData{
			expires,
		}, nil
	}
	return nil, err

}

func CheckToken(c *fiber.Ctx) (bool, error) {

	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(c)

	if err != nil {
		return false, err
	}
	if now > claims.Expires {
		return false, nil
	}

	return true, nil
}

func ExtractToken(c *fiber.Ctx) string {

	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(GetValue("JWT_SECRET_KEY")), nil
}
