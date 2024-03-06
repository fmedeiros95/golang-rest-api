package utils

import (
	"context"
	"fmedeiros95/golang-rest-api/models"
)

type contextKey string

const userKey contextKey = "user"

func SetUserInContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) *models.User {
	if user, ok := ctx.Value(userKey).(*models.User); ok {
		return user
	}
	return nil
}
