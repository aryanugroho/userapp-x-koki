package usecase

import (
	"context"

	"github.com/aryanugroho/userapp-x-koki/common/jwt"
	"github.com/aryanugroho/userapp-x-koki/infrastructure"
)

type Application interface {
	NewUser()
	UpdateUser()
	GetUser()
	DeleteUser()
	GetUsers()

	Login(ctx context.Context, request LoginRequest) (*Authenticated, error)
	Logout(ctx context.Context, request LogoutRequest) error
	RefreshToken(ctx context.Context, request RefreshRequest) (*Authenticated, error)
}

type UserApp struct {
	store        infrastructure.Store
	tokenManager jwt.TokenManager
}
