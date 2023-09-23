package app

import (
	"Store/internal/oas"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-faster/errors"
)

const secretKey = "YOUR_SECRET_KEY"

type MySecurityHandler struct{}

func (m MySecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t oas.BearerAuth) (context.Context, error) {
	// Parse the JWT token from the BearerAuth
	tokenString := t.Token

	// Parse the token
	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if parseErr != nil {
		return ctx, parseErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Here you can extract claims if needed and set them into the context
		// For example:
		userID, exists := claims["user_id"].(string)
		if !exists {
			return ctx, errors.New("user_id claim not found")
		}

		// You can store the userID or other claims in the context
		ctx = context.WithValue(ctx, "userID", userID)
		return ctx, nil
	} else {
		return ctx, errors.New("invalid token")
	}
}
