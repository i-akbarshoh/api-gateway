package user

import (
	"context"

	"github.com/i-akbarshoh/api-gateway/internal/pkg/proto"
	"github.com/i-akbarshoh/api-gateway/internal/entity"
)

type user struct {
	client proto.AuthClient
}

func NewUsecase(ac *proto.AuthClient) *user {
	return &user{client: *ac}
}

func (uc *user) SignUp(ctx context.Context, user entity.User) error {
	if _, err := uc.client.SignUp(ctx, &proto.SignUpModel{
		Id: user.ID,
		Name: user.Name,
		Age: user.Age,
		Email: user.Email,
		Password: user.Password,
		Role: user.Role,
	}); err != nil {
		return err
	}

	return nil
}
func (uc *user) Login(ctx context.Context, login entity.Login) error {
	uc.client.Login(ctx, &proto.LoginModel{
		Email: login.Email,
		Password: login.Password,
	})
	return nil
}