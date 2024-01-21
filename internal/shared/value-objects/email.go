package valueobjects

import "github.com/read-manager/read-manager-main-api/internal/shared/validator"

type email struct {
	value string
}

func NewEmail(v *validator.Validator, value string) *email {
	key := "email"
	v.Check(value != "", key, "must be provided")
	v.Check(validator.Matches(value, validator.EmailRX), key, "must be a valid email")
	if _, exists := v.Errors[key]; exists {
		return nil
	}
	return &email{ value: value }
}

func (e *email) Get() string {
    return e.value
}
