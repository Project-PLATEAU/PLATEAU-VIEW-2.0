package user

import (
	"bytes"
	"errors"
	"unicode"

	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"golang.org/x/crypto/bcrypt"
)

var (
	DefaultPasswordEncoder PasswordEncoder = &BcryptPasswordEncoder{}
	ErrEncodingPassword                    = rerror.NewE(i18n.T("encoding password"))
	ErrInvalidPassword                     = rerror.NewE(i18n.T("invalid password"))
	ErrPasswordLength                      = rerror.NewE(i18n.T("password at least 8 characters"))
	ErrPasswordUpper                       = rerror.NewE(i18n.T("password should have upper case letters"))
	ErrPasswordLower                       = rerror.NewE(i18n.T("password should have lower case letters"))
	ErrPasswordNumber                      = rerror.NewE(i18n.T("password should have numbers"))
)

type PasswordEncoder interface {
	Encode(string) ([]byte, error)
	Verify(string, []byte) (bool, error)
}

type BcryptPasswordEncoder struct{}

func (BcryptPasswordEncoder) Encode(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 14)
}

func (BcryptPasswordEncoder) Verify(s string, p []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p, []byte(s))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

type NoopPasswordEncoder struct{}

func (m NoopPasswordEncoder) Encode(pass string) ([]byte, error) {
	return []byte(pass), nil
}

func (m NoopPasswordEncoder) Verify(s string, p []byte) (bool, error) {
	return bytes.Equal([]byte(s), []byte(p)), nil
}

type MockPasswordEncoder struct{ Mock []byte }

func (m MockPasswordEncoder) Encode(pass string) ([]byte, error) {
	return append(m.Mock[:0:0], m.Mock...), nil
}

func (m MockPasswordEncoder) Verify(s string, p []byte) (bool, error) {
	return bytes.Equal(m.Mock, []byte(s)), nil
}

type EncodedPassword []byte

func NewEncodedPassword(pass string) (EncodedPassword, error) {
	if err := ValidatePasswordFormat(pass); err != nil {
		return nil, err
	}
	got, err := DefaultPasswordEncoder.Encode(pass)
	if err != nil {
		return nil, ErrEncodingPassword
	}
	return got, nil
}

func MustEncodedPassword(pass string) EncodedPassword {
	p, err := NewEncodedPassword(pass)
	if err != nil {
		panic(err)
	}
	return p
}

func (p EncodedPassword) Clone() EncodedPassword {
	if p == nil {
		return nil
	}
	return append(p[:0:0], p...)
}

func (p EncodedPassword) Verify(toVerify string) (bool, error) {
	if len(toVerify) == 0 || len(p) == 0 {
		return false, nil
	}
	return DefaultPasswordEncoder.Verify(toVerify, p)
}

func ValidatePasswordFormat(pass string) error {
	var hasNum, hasUpper, hasLower bool
	for _, c := range pass {
		switch {
		case unicode.IsNumber(c):
			hasNum = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c) || c == ' ':
			hasLower = true
		}
	}
	if len(pass) < 8 {
		return ErrPasswordLength
	}
	if !hasLower {
		return ErrPasswordLower
	}
	if !hasUpper {
		return ErrPasswordUpper
	}
	if !hasNum {
		return ErrPasswordNumber
	}
	return nil
}
