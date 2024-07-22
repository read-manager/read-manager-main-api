package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"


type userName struct {
	value string
}

func NewUserName(v *validator.Validator, value string) *userName {
	key := "user-name"
	v.Check(value != "", key, "must be provided")
	v.Check(len(value) >= 2, key, "must contain at least 2 bytes long")
	v.Check(len(value) <= 20, key, "must not be more than 20 bytes long")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &userName{ value: value }
}

func (un *userName) Get() string {
    return un.value
}
