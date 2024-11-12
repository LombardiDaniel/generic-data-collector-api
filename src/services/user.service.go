package services

import (
	"context"

	"github.com/LombardiDaniel/generic-data-collector-api/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, username string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	NoUsersRegistered(ctx context.Context) (bool, error)
}