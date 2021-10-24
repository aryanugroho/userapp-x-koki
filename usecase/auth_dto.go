package usecase

import "errors"

type LoginDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (request *LoginDTO) validate() error {
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

type LogoutDTO struct {
	AccessToken string `json:"access_token"`
}

func (request *LogoutDTO) validate() error {
	if request.AccessToken == "" {
		return errors.New("access token is required")
	}

	return nil
}

type RefreshDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (request *RefreshDTO) validate() error {
	if request.RefreshToken == "" {
		return errors.New("refresh token is required")
	}

	return nil
}
