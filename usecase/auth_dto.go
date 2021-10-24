package usecase

import "errors"

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (request *LoginRequest) validate() error {
	if request.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if request.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type Authenticated struct {
	AccessToken  string
	RefreshToken string
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

func (request *LogoutRequest) validate() error {
	if request.AccessToken == "" {
		return errors.New("access token is required")
	}

	return nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (request *RefreshRequest) validate() error {
	if request.RefreshToken == "" {
		return errors.New("refresh token is required")
	}

	return nil
}
