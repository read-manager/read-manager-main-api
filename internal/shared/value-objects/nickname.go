package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"

type nickname struct {
	value string
}

func NewNickname(v *validator.Validator, value string) *nickname {
	key := "nickname"
	v.Check(value != "", key, "must be provided")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &nickname{ value: value }
}

func (e *nickname) Get() string {
    return e.value
}
