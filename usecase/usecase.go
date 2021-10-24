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

	Login(ctx context.Context, request LoginDTO) (*Authenticated, error)
	Logout(ctx context.Context, request LogoutDTO) error
	RefreshToken(ctx context.Context, request RefreshDTO) (*Authenticated, error)
}

type UserApp struct {
	store        infrastructure.Store
	tokenManager jwt.TokenManager
}
