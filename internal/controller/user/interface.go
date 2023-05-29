package user

import (
	"context"

	"github.com/i-akbarshoh/api-gateway/internal/entity"
)

type Usecase interface {
	SignUp(context.Context, entity.User) error
	Login(context.Context, entity.Login) error
}