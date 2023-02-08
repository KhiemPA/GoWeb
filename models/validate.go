package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

var (
	ErrRequiredFirstName = errors.New("Firstname required")
	ErrRequiredLastName = errors.New("Lastname required")
	ErrRequiredEmail = errors.New("Email required")
	ErrRequiredPassword = errors.New("Password required")
	ErrInvalidEmail = errors.New("Email invalid")
	ErrMaxLimit = errors.New("Many char!!!")
	ErrDuplicateKeyEmail = errors.New("Email is used")
)

func IsEmpty( str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func Max(str string, lim int) bool {
	if len(str) <= lim {
		return true
	}
	return false
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}


func ValidateLimitField(user User) (User, error) {
	if !Max(user.FirstName, 15) || !Max(user.LastName, 20) || !Max(user.Email, 40) || !Max(user.Password, 100) {
		return user, ErrMaxLimit
	}
	return user, nil
}

func UniqueEmail(email string) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := " select count(email) from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()
	var count int64
	if rs.Next(){
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	if count > 0 {
		return false, ErrDuplicateKeyEmail
	}
	return true, nil
}

func ValidateNewUser(u User) (User, error) {
	_, err := UniqueEmail(u.Email)
	if err != nil {
		return User{}, err
	}
	u, err = ValidateLimitField(u)
	if err != nil {
		return u, err
	}
	u.FirstName = Trim(u.FirstName)
	u.LastName = Trim(u.LastName)
	u.Email = Trim(strings.ToLower(u.Email))
	if IsEmpty(u.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(u.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(u.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(u.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(u.Password) {
		return User{}, ErrRequiredPassword
	}
	return u, nil
}