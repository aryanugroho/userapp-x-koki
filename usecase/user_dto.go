package usecase

import (
	"errors"
)

var (
	ErrFieldIDIsRequired          = errors.New("id is required")
	ErrFieldNameIsRequired        = errors.New("name is required")
	ErrFieldPhoneNumberIsRequired = errors.New("phone_number is required")
	ErrFieldEmailIsRequired       = errors.New("email is required")
	ErrFieldPasswordIsRequired    = errors.New("password is required")
)

type UserCreateDTO struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
}

func (userDTO *UserCreateDTO) validate() error {
	if userDTO.Name == "" {
		return ErrFieldNameIsRequired
	}
	if userDTO.PhoneNumber == "" {
		return ErrFieldPhoneNumberIsRequired
	}
	if userDTO.Email == "" {
		return ErrFieldEmailIsRequired
	}
	if userDTO.Password == "" {
		return ErrFieldPasswordIsRequired
	}
	return nil
}

type UserUpdateDTO struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
}

func (userDTO *UserUpdateDTO) validate() error {
	if userDTO.ID == "" {
		return ErrFieldIDIsRequired
	}
	if userDTO.Name == "" {
		return ErrFieldNameIsRequired
	}
	if userDTO.PhoneNumber == "" {
		return ErrFieldPhoneNumberIsRequired
	}
	if userDTO.Email == "" {
		return ErrFieldEmailIsRequired
	}
	if userDTO.Password == "" {
		return ErrFieldPasswordIsRequired
	}
	return nil
}

type GetUsersDTO struct {
	Total   int64
	PageNum int
	Size    int
	Data    interface{}
}
