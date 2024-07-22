package users

import (
	"time"

	"github.com/google/uuid"

	"github.com/read-manager/read-manager-main-api/internal/shared/validator"
	valueobjects "github.com/read-manager/read-manager-main-api/internal/shared/value-objects"
)


type User struct {
	Id         uuid.UUID
	Email      string
	Password   valueobjects.Password
	Name       string
	Nickname   string
	ReadPoints int64
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewUser(email string, password string, name string, nickname string) (*User, map[string]string) {
	v := validator.New()
	emailVO := valueobjects.NewEmail(v, email)
    passwordVO := valueobjects.NewPassword(v, password)
    userNameVO := valueobjects.NewUserName(v, name)
    nicknameVO := valueobjects.NewNickname(v, nickname)
	if !v.Valid() {
		return &User{}, v.Errors
	}
    return &User{
        Id: uuid.New(),
        Email: emailVO.Get(),
        Password: *passwordVO,
        Name: userNameVO.Get(),
        Nickname: nicknameVO.Get(),
    }, nil
}
