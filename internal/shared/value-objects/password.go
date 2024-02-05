package valueobjects

import (
	"errors"

	"github.com/read-manager/read-manager-main-api/internal/shared/validator"
	"golang.org/x/crypto/bcrypt"
)

type password struct {
	plaintext *string
	hash      []byte
}

func NewPassword(v *validator.Validator, value string) *password {
	key := "password"
	v.Check(value != "", key, "must be provided")
	v.Check(len(value) >= 8, key, "must contain at least 8 bytes long")
	v.Check(len(value) <= 72, key, "must not be more than 72 bytes long")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 12)
	if err != nil {
		v.AddError("password", "invalid password")
		return nil
	}
	return &password{plaintext: &value, hash: hash}
}

func (p *password) GetHash() []byte {
	return p.hash
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
