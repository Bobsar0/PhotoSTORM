package context

import (
	"context"

	"github.com/bobsar0/PhotoSTORM/models"
)

type privateKey string

const (
	userKey privateKey = "user"
)

//WithUser accepts an existing context and a user, and then returns a new context with that user set as a value.
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

//User looks up a user from a given context.
func User(ctx context.Context) *models.User {
	if temp := ctx.Value(userKey); temp != nil { //verify that a user was previously stored in the context
		if user, ok := temp.(*models.User); ok {
			return user
		}
	}
	return nil
}
