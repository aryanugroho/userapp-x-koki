package usecase

import (
	"context"
	"errors"

	"github.com/aryanugroho/userapp-x-koki/common/hash"
	"github.com/aryanugroho/userapp-x-koki/common/log"
	"github.com/aryanugroho/userapp-x-koki/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func (app *UserApp) NewUser(ctx context.Context, request *UserCreateDTO) (*domain.User, error) {
	// validate
	err := request.validate()
	if err != nil {
		log.Error(ctx, "failed to validate", err)
		return nil, err
	}

	// encrypt the password
	encryptedPassword, err := hash.GetProvider().Hash(request.Password)
	if err != nil {
		log.Error(ctx, "failed to Hash", err)
		return nil, err
	}

	// save the data
	savedUser, err := app.Store.User().Add(ctx, &domain.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    encryptedPassword,
	})
	if err != nil {
		log.Error(ctx, "failed to Add", err)
		return nil, err
	}

	return savedUser, nil
}

func (app *UserApp) UpdateUser(ctx context.Context, request *UserUpdateDTO) (*domain.User, error) {
	// validate
	err := request.validate()
	if err != nil {
		log.Error(ctx, "failed to validate", err)
		return nil, err
	}

	// encrypt the password
	encryptedPassword, err := hash.GetProvider().Hash(request.Password)
	if err != nil {
		log.Error(ctx, "failed to Hash", err)
		return nil, err
	}

	// save the data
	savedUser, err := app.Store.User().Update(ctx, &domain.User{
		ID:          request.ID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    encryptedPassword,
	})
	if err != nil {
		log.Error(ctx, "failed to Add", err)
		return nil, err
	}

	return savedUser, nil
}

func (app *UserApp) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// validate
	if id == "" {
		return nil, ErrUserNotFound
	}

	user, err := app.Store.User().Get(ctx, id)
	if err != nil {
		log.Error(ctx, "failed to Get", err)
		return nil, err
	}

	return user, nil
}

func (app *UserApp) DeleteUser(ctx context.Context, id string) error {
	// validate
	if id == "" {
		return ErrUserNotFound
	}

	// delete the user
	err := app.Store.User().Delete(ctx, id)
	if err != nil {
		log.Error(ctx, "failed to Delete", err)
		return err
	}

	return nil
}

func (app *UserApp) GetUsers(ctx context.Context, params map[string]interface{}, pageNum, size int) (*GetUsersDTO, error) {
	// validate
	if pageNum < 1 || size < 1 {
		return nil, ErrUserNotFound
	}

	users, err := app.Store.User().GetAll(ctx, params, pageNum, size)
	if err != nil {
		log.Error(ctx, "failed to GetAll", err)
		return nil, err
	}

	count, err := app.Store.User().CountAll(ctx, params)
	if err != nil {
		log.Error(ctx, "failed to CountAll", err)
		return nil, err
	}

	return &GetUsersDTO{
		Total:   count,
		PageNum: pageNum,
		Size:    size,
		Data:    users,
	}, nil
}
