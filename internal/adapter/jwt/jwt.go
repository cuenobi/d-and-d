package jwt

import (
	"strings"
	"time"

	"d-and-d/config"
	"d-and-d/internal/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	SecretKey string
	Exp       int
}

func NewJwtToken(cfg *config.JwtConfig) *Jwt {
	return &Jwt{
		SecretKey: cfg.Secret,
		Exp:       cfg.Exp,
	}
}

type CustomClaims struct {
	Username string
	Role     uint
	jwt.RegisteredClaims
}

func (j *Jwt) Generate(username string, role uint) string {
	claims := CustomClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Exp) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return ""
	}

	return ss
}

func (j *Jwt) Validate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		username := claims.Username
		role := claims.Role

		c.Locals("username", username)
		c.Locals("role", role)
	}

	return c.Next()
}

func (j *Jwt) ValidateAdmin(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		username := claims.Username
		role := claims.Role

		if constant.Role(role) != constant.Admin {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid permission",
			})
		}

		c.Locals("username", username)
		c.Locals("role", role)
	}

	return c.Next()
}
