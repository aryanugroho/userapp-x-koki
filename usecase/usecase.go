package usecase

import (
	"context"

	"github.com/aryanugroho/userapp-x-koki/common/jwt"
	"github.com/aryanugroho/userapp-x-koki/domain"
	"github.com/aryanugroho/userapp-x-koki/infrastructure"
)

type Application interface {
	NewUser(ctx context.Context, request *UserCreateDTO) (*domain.User, error)
	UpdateUser(ctx context.Context, request *UserUpdateDTO) (*domain.User, error)
	GetUser(ctx context.Context, id string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, params map[string]interface{}, pageNum, size int) ([]*domain.User, error)

	Login(ctx context.Context, request LoginDTO) (*Authenticated, error)
	Logout(ctx context.Context, request LogoutDTO) error
	RefreshToken(ctx context.Context, request RefreshDTO) (*Authenticated, error)
}

type UserApp struct {
	Store        infrastructure.Store
	TokenManager *jwt.TokenManager
}
