package usecase

import (
	"context"
	"errors"

	"github.com/aryanugroho/userapp-x-koki/common/hash"
	"github.com/aryanugroho/userapp-x-koki/common/jwt"
	"github.com/aryanugroho/userapp-x-koki/common/log"
)

var (
	ErrCredentialNotValid = errors.New("credential not valid")
	ErrInternal           = errors.New("internal error")
)

func (app *UserApp) Login(ctx context.Context, request LoginRequest) (*Authenticated, error) {
	// validate
	err := request.validate()
	if err != nil {
		log.Error(ctx, "failed to validate", err)
		return nil, err
	}

	// lookup user by phone
	user, err := app.store.User().FindByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		log.Error(ctx, "failed to FindByPhoneNumber", err)
		return nil, ErrInternal
	}

	// compare cred
	if !hash.GetProvider().Compare(request.Password, user.Password) {
		log.Error(ctx, "failed to Compare", err)
		return nil, ErrCredentialNotValid
	}

	// generate token
	tokenData := jwt.UserForToken{
		Email:       user.Email,
		UserID:      user.ID,
		PhoneNumber: user.PhoneNumber,
	}
	accessToken, err := app.tokenManager.GenerateAccessToken(&tokenData)
	if err != nil {
		log.Error(ctx, "failed to GenerateAccessToken", err)
		return nil, ErrInternal
	}
	refreshToken, err := app.tokenManager.GenerateRefreshToken(&tokenData)
	if err != nil {
		log.Error(ctx, "failed to GenerateRefreshToken", err)
		return nil, ErrInternal
	}

	return &Authenticated{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (app *UserApp) Logout(ctx context.Context, request LogoutRequest) error {
	return nil
}

func (app *UserApp) RefreshToken(ctx context.Context, request RefreshRequest) (*Authenticated, error) {
	return nil, nil
}
