package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/AmongAster/golang-todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullname string,
	phonenamber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullname,
		PhoneNumber: phonenamber,
	}
}

func NewUserUninitializer(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializerID,
		UninitializerVersionID,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLangth := len([]rune(u.FullName))
	if fullNameLangth < 3 || fullNameLangth > 100 {
		return fmt.Errorf(
			"invalid fulName len %d: %w",
			fullNameLangth,
			core_errors.ErrinvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"ivolid Phonenumber len %d: %w",
				phoneNumberLen,
				core_errors.ErrinvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid Phonenumber format: %w", core_errors.ErrinvalidArgument)
		}
	}

	return nil

}

// GET /users/{id}
