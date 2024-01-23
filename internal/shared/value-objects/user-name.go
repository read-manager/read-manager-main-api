package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"


type userName struct {
	value string
}

func NewUserName(v *validator.Validator, value string) *userName {
	key := "user-name"
	v.Check(value != "", key, "must be provided")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &userName{ value: value }
}

func (un *userName) Get() string {
    return un.value
}
