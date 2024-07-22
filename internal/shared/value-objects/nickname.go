package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"

type nickname struct {
	value string
}

func NewNickname(v *validator.Validator, value string) *nickname {
	key := "nickname"
	v.Check(value != "", key, "must be provided")
	v.Check(len(value) >= 2, key, "must contain at least 2 bytes long")
	v.Check(len(value) <= 10, key, "must not be more than 10 bytes long")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &nickname{ value: value }
}

func (e *nickname) Get() string {
    return e.value
}
