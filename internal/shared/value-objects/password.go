package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"


type password struct {
	value string
}

func NewPassword(v *validator.Validator, value string) *password {
	key := "password"
	v.Check(value != "", key, "must be provided")
	v.Check(len(value) > 8, key, "must contain at least 8 bytes long")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &password{ value: value }
}

func (p *password) Get() string {
    return p.value
}
