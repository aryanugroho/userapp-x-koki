package infrastructure

import (
	"context"

	"github.com/aryanugroho/userapp-x-koki/domain"
)

type Store interface {
	User() UserStore
}

type UserStore interface {
	Add(ctx context.Context, model *domain.User) (*domain.User, error)
	Update(ctx context.Context, model *domain.User) (*domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	GetAll(ctx context.Context, params map[string]interface{}, pageNum, size int) ([]*domain.User, error)
	CountAll(ctx context.Context, params map[string]interface{}) (int64, error)
}
