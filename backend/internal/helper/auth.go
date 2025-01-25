package helper

import (
	"errors"
	"fmt"
	"nxg/internal/domain"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a *Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) <= 3 {
		return "", errors.New("password must be at least 6 characters long")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", errors.New("failed to generate password hashed")
	}
	return string(hashedPassword), nil
}

func (a *Auth) GenerateToken(id uint, email string, role int) (string, error) {
	if id == 0 || email == "" || role == 0 {
		return "", errors.New("required inputs are missing for generating the tokens")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("failed to sign the token")
	}

	return tokenString, nil
}

func (a *Auth) VerifyPassword(p string, hp string) error {
	if len(p) < 4 {
		return errors.New("password must be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
	if err != nil {
		return errors.New("password does not match")
	}
	return nil
}

func (a *Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	tokenString := tokenArr[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}
		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = int(claims["role"].(float64))
		return user, nil
	}
	return domain.User{}, nil
}

func (a *Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]
	fmt.Println("authheader data: ", authHeader)
	user, err := a.VerifyToken(strings.Join(authHeader, ""))
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}
}

func (a *Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}
