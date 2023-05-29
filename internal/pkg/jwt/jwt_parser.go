package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/config"
)

type TokenMetadata struct {
	UserID      uuid.UUID
	Credentials map[string]string
	Expires     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		expires := claims["expires"].(float64)

		credentials := map[string]string{
			"role": claims["role"].(string),
		}

		return &TokenMetadata{
			UserID:      userID,
			Credentials: credentials,
			Expires:     int64(expires),
		}, nil
	}

	return nil, err
}

func extractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")

	return bearToken
}

func verifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.C.JWT.SigningKey), nil
}
